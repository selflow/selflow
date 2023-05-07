package workflow

import (
	"context"
	"errors"
	"fmt"
	"github.com/selflow/selflow/pkg/sflog"
	"log"
	"reflect"
	"sync"
)

type SimpleWorkflow struct {
	steps        []Step
	dependencies map[Step][]Step
}

type Workflow interface {
	Init() error
	AddStep(step Step, dependencies []Step) error
	Execute(ctx context.Context) (map[string]map[string]string, error)
	Equals(s2 Workflow) bool
}

var stepOutputContextKey struct{}

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
	if dependencies, ok := s.dependencies[currentStep]; ok {
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
func NewWorkflow(stepCount uint) Workflow {
	return &SimpleWorkflow{
		steps:        make([]Step, 0, stepCount),
		dependencies: make(map[Step][]Step),
	}
}

func (s *SimpleWorkflow) getNextSteps() []Step {
	nextSteps := make([]Step, 0, len(s.steps))
	for _, step := range s.steps {
		if !step.GetStatus().IsFinished() && step.GetStatus() != RUNNING && areRequirementsFullFilled(step, s.dependencies) {
			nextSteps = append(nextSteps, step)
		}
	}
	return nextSteps
}

func (s *SimpleWorkflow) executeStep(ctx context.Context, step Step) {
	logger := sflog.LoggerFromContext(ctx)
	requirementsOutputs := s.getRequirementsOutputs(step)
	stepContext := context.WithValue(ctx, stepOutputContextKey, requirementsOutputs)
	_, err := step.Execute(stepContext)
	if err != nil {
		logger.Warn("step ended with an error", "step-id", step.GetId(), "error", err)
	}
}

func (s *SimpleWorkflow) cancelNextSteps(lastStep Step, closingSteps chan Step) error {
	var err error

	concernedSteps := getStepThatRequires(lastStep, s.dependencies)

	for _, step := range concernedSteps {
		if step.GetStatus().IsCancellable() {
			err = errors.Join(step.Cancel())
			closingSteps <- step
		}
	}
	return err
}

func (s *SimpleWorkflow) getRequirementsOutputs(step Step) map[string]map[string]string {
	res := make(map[string]map[string]string)
	stepDependencies := s.dependencies[step]

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

func (s *SimpleWorkflow) Execute(ctx context.Context) (map[string]map[string]string, error) {
	logger := sflog.LoggerFromContext(ctx)
	closingSteps := make(chan Step, len(s.steps))

	var err error

	activeSteps := &sync.WaitGroup{}

	s.startNextSteps(ctx, activeSteps, closingSteps)

	for s.hasUnfinishedSteps() {
		select {
		case <-ctx.Done():
			err = errors.Join(err, s.cancelRemainingSteps())
			close(closingSteps)

		case step := <-closingSteps:
			logger.Info("step terminated", "step-id", step.GetId(), "status", step.GetStatus().GetName())
			// A step as ended
			if step.GetStatus() == ERROR || step.GetStatus() == CANCELLED {
				err = errors.Join(err, s.cancelNextSteps(step, closingSteps))
			} else {
				s.startNextSteps(ctx, activeSteps, closingSteps)
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
	logger := sflog.LoggerFromContext(ctx)
	nextSteps := s.getNextSteps()
	for _, step := range nextSteps {
		activeSteps.Add(1)

		go func(step Step) {
			logger.Info("step started", "step-id", step.GetId())
			s.executeStep(ctx, step)
			closingSteps <- step
			activeSteps.Done()
		}(step)
	}
}

func (s *SimpleWorkflow) debug() {
	for _, step := range s.steps {
		log.Printf("[DEBUG]: %v : %v", step.GetId(), step.GetStatus().GetName())
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
	s.dependencies[wrappedStep] = dependencies
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
