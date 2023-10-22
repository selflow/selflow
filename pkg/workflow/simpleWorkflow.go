package workflow

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"reflect"
	"sync"
)

type SimpleWorkflow struct {
	steps        []Step
	Dependencies map[Step][]Step
	StateCh      chan map[string]Status
}

type Workflow interface {
	Init() error
	AddStep(step Step, dependencies []Step) error
	Execute(ctx context.Context) (map[string]map[string]string, error)
	Equals(s2 Workflow) bool
}

var StepOutputContextKey struct{}

// hasCycle check if the workflow has cycle starting from currentStep.
// visited contains the list of already visited steps.
// secureCycles contains the list of all steps that have already been checked and can be skipped.
func (s *SimpleWorkflow) hasCycle(visited []Step, secureCycles []Step, currentStep Step) (bool, []Step) {
	if sliceContainsStep(secureCycles, currentStep) {
		// The step has already been checked and has no cycle
		return false, []Step{}
	}
	if sliceContainsStep(visited, currentStep) {
		// The step has already been visited, this is a cycle
		return true, []Step{}
	}
	visited = append(visited, currentStep)

	currentStepSecureCycles := make([]Step, 0, len(s.steps))

	// Check for each step if its dependencies contains cycles
	if dependencies, ok := s.Dependencies[currentStep]; ok {
		for _, dependency := range dependencies {
			if hasCycle, dependencySecureCycles := s.hasCycle(visited, secureCycles, dependency); !hasCycle {
				currentStepSecureCycles = append(currentStepSecureCycles, dependencySecureCycles...)
			} else {
				// A dependency has a cycle, return true and no secure steps
				return true, []Step{}
			}
		}
	}
	// this dependency is safe
	currentStepSecureCycles = append(currentStepSecureCycles, currentStep)
	return false, currentStepSecureCycles
}

func (s *SimpleWorkflow) Init() error {
	visitedSteps := make([]Step, 0, len(s.steps))
	secureSteps := make([]Step, 0, len(s.steps))

	for _, currentStep := range s.steps {
		if hasCycles, stepDependency := s.hasCycle(visitedSteps, secureSteps, currentStep); !hasCycles {
			secureSteps = append(secureSteps, stepDependency...)
		} else {
			return fmt.Errorf("cycle detected")
		}
	}

	return nil
}

// NewWorkflow creates a SimpleWorkflow instance
func NewWorkflow(stepCount uint) *SimpleWorkflow {
	return &SimpleWorkflow{
		steps:        make([]Step, 0, stepCount),
		Dependencies: make(map[Step][]Step),
		StateCh:      make(chan map[string]Status),
	}
}

func (s *SimpleWorkflow) getNextSteps() []Step {
	nextSteps := make([]Step, 0, len(s.steps))
	for _, step := range s.steps {
		if !step.GetStatus().IsFinished() && step.GetStatus() != RUNNING && areRequirementsFullFilled(step, s.Dependencies) {
			nextSteps = append(nextSteps, step)
		}
	}
	return nextSteps
}

func (s *SimpleWorkflow) executeStep(ctx context.Context, step Step) {
	requirementsOutputs := s.getRequirementsOutputs(step)
	stepContext := context.WithValue(ctx, StepOutputContextKey, requirementsOutputs)
	_, err := step.Execute(stepContext)
	if err != nil {
		slog.WarnContext(ctx, "Step ended with an error", "stepId", step.GetId(), "error", err)
	}
}

func (s *SimpleWorkflow) cancelNextSteps(lastStep Step, closingSteps chan Step) error {
	var err error

	concernedSteps := getStepThatRequires(lastStep, s.Dependencies)

	for _, step := range concernedSteps {
		if step.GetStatus().IsCancellable() {
			err = errors.Join(step.Cancel())
			closingSteps <- step

			err = errors.Join(s.cancelNextSteps(step, closingSteps))
		}
	}
	return err
}

func (s *SimpleWorkflow) getRequirementsOutputs(step Step) map[string]map[string]string {
	res := make(map[string]map[string]string)
	stepDependencies := s.Dependencies[step]

	for _, dependency := range stepDependencies {
		res[dependency.GetId()] = dependency.GetOutput()
		res = mergeStringStringStringMaps(res, s.getRequirementsOutputs(dependency))
	}

	return res
}

func (s *SimpleWorkflow) getOutput() map[string]map[string]string {
	output := make(map[string]map[string]string)
	for _, step := range s.steps {
		output[step.GetId()] = step.GetOutput()
	}
	return output
}

func (s *SimpleWorkflow) hasUnfinishedSteps() bool {
	for _, step := range s.steps {
		if !step.GetStatus().IsFinished() {
			return true
		}
	}
	return false
}

func (s *SimpleWorkflow) updateState() {
	state := map[string]Status{}
	for _, step := range s.steps {
		state[step.GetId()] = step.GetStatus()
	}

	s.StateCh <- state
}

func shouldCancelNextSteps(stepStatus Status) bool {
	return stepStatus.GetCode() == ERROR.GetCode() || stepStatus.GetCode() == CANCELLED.GetCode()
}

func (s *SimpleWorkflow) Execute(ctx context.Context) (map[string]map[string]string, error) {
	closingSteps := make(chan Step, len(s.steps))
	activeSteps := &sync.WaitGroup{}

	defer close(s.StateCh)

	for s.hasUnfinishedSteps() {
		s.startNextSteps(ctx, activeSteps, closingSteps)

		select {
		case <-ctx.Done():
			// The context has been closed
			slog.Debug("Context has expired")

			err := s.cancelRemainingSteps()
			if err != nil {
				slog.ErrorContext(ctx, "Cancel error", "error", err)
			}
			close(closingSteps)

		case step := <-closingSteps:
			// A step as ended

			s.updateState()
			slog.InfoContext(ctx, "Step terminated", "stepId", step.GetId(), "stepStatus", step.GetStatus().GetName())
			if shouldCancelNextSteps(step.GetStatus()) {
				err := s.cancelNextSteps(step, closingSteps)
				if err != nil {
					slog.ErrorContext(ctx, "Cancel error", "error", err)
				}
			}
		}
	}

	activeSteps.Wait()

	return s.getOutput(), nil
}

func (s *SimpleWorkflow) cancelRemainingSteps() error {
	var err error
	for _, step := range s.steps {
		if !step.GetStatus().IsFinished() && step.GetStatus().IsCancellable() {
			err = errors.Join(step.Cancel())
		}
	}
	return err
}

func (s *SimpleWorkflow) startNextSteps(ctx context.Context, activeSteps *sync.WaitGroup, closingSteps chan Step) {
	nextSteps := s.getNextSteps()
	for _, step := range nextSteps {
		activeSteps.Add(1)

		go func(step Step) {
			slog.InfoContext(ctx, "Step started", "stepId", step.GetId())
			s.executeStep(ctx, step)
			closingSteps <- step
			activeSteps.Done()
		}(step)
	}
}

func (s *SimpleWorkflow) AddStep(step Step, dependencies []Step) error {
	wrappedStep := wrapStep(step)
	for _, previousStep := range s.steps {
		if previousStep.GetId() == wrappedStep.GetId() {
			return fmt.Errorf("step [%s] is already present in workflow", step.GetId())
		}
	}

	s.steps = append(s.steps, wrappedStep)
	s.Dependencies[wrappedStep] = dependencies
	return nil
}

func (s *SimpleWorkflow) Equals(s2 Workflow) bool {
	sw2, ok := s2.(*SimpleWorkflow)
	if !ok {
		return false
	}

BaseStep:
	for _, step := range s.steps {
		for _, step2 := range sw2.steps {
			if reflect.DeepEqual(step.GetId(), step2.GetId()) {
				continue BaseStep
			}
		}
		return false
	}

	return true
}

var _ Workflow = &SimpleWorkflow{}
