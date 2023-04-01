package wfbuilder

import (
	"fmt"
	"github.com/selflow/selflow/internal/config"
	"github.com/selflow/selflow/pkg/workflow"
)

// StepMapper represents a function that can create a step from a step definition
type StepMapper = func(string, config.StepDefinition) (workflow.Step, error)

// Builder is a structure that creates selflow workflows
type Builder struct {
	StepBuilderMap map[string]StepMapper
}

// mapDefinitionToStep creates a step from the given step definition. If the step couldn't be created, an error is returned
func (b Builder) mapDefinitionToStep(stepId string, stepDefinition config.StepDefinition) (step workflow.Step, err error) {
	for stepType, stepMapper := range b.StepBuilderMap {
		if stepDefinition.Kind == stepType {
			step, err = stepMapper(stepId, stepDefinition)
			if err == nil {
				return step, nil
			}
		}
	}
	return
}

// mapStepsDefinitionsToSteps cast a map of step definition to a map of step.
// Throws an error if a step couldn't be cast
func (b Builder) mapStepsDefinitionsToSteps(definitions config.Steps) (map[string]workflow.Step, error) {
	steps := map[string]workflow.Step{}

	for stepId, stepDefinition := range definitions {
		step, err := b.mapDefinitionToStep(stepId, stepDefinition)
		if err != nil {
			return nil, err
		}

		steps[stepId] = step
	}

	return steps, nil
}

// BuildWorkflow creates a workflow from a config
func (b Builder) BuildWorkflow(parsedConfig *config.Flow) (workflow.Workflow, error) {

	parsedSteps, err := b.mapStepsDefinitionsToSteps(parsedConfig.Workflow.Steps)
	if err != nil {
		return nil, err
	}

	wf := workflow.NewWorkflow(uint(len(parsedConfig.Workflow.Steps)))

	for stepId, stepDefinition := range parsedConfig.Workflow.Steps {
		requirements := make([]workflow.Step, len(stepDefinition.Needs))

		for requirementIndex, requiredStepId := range stepDefinition.Needs {
			requiredStep, ok := parsedSteps[requiredStepId]
			if !ok {
				return nil, fmt.Errorf("missing dependency %s for step %s", requiredStepId, stepId)
			}
			requirements[requirementIndex] = requiredStep
		}

		err = wf.AddStep(parsedSteps[stepId], requirements)
		if err != nil {
			return nil, err
		}
	}
	return wf, err
}
