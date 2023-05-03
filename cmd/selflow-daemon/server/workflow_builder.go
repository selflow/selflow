package server

import (
	"errors"
	"fmt"
	"github.com/selflow/selflow/internal/config"
	"github.com/selflow/selflow/pkg/workflow"
)

var UnknownStepKindError = errors.New("unknown step kind")

type WorkflowBuilder struct {
	stepMappers []StepMapper
}

func (b WorkflowBuilder) mapDefinitionToStep(stepId string, stepDefinition config.StepDefinition) (workflow.Step, error) {

	for _, stepMapper := range b.stepMappers {
		step, err := stepMapper.MapStep(stepId, stepDefinition)
		if err == nil {
			return step, nil
		}
	}

	return nil, UnknownStepKindError
}

func (b WorkflowBuilder) mapStepsDefinitionsToSteps(definitions config.Steps) (map[string]workflow.Step, error) {
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

func (b WorkflowBuilder) BuildWorkflow(parsedConfig *config.Flow) (workflow.Workflow, error) {

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
