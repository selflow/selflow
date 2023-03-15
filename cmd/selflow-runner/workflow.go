package main

import (
	"fmt"
	"github.com/selflow/selflow/internal/config"
	dockerStep "github.com/selflow/selflow/pkg/docker-step"
	"github.com/selflow/selflow/pkg/workflow"
)

func mapDefinitionToStep(stepId string, stepDefinition config.StepDefinition) (step workflow.Step, err error) {
	if stepDefinition.Kind == "docker" {
		step, err = dockerStep.NewDockerStep(stepId, stepDefinition)
	}
	return
}

func mapStepsDefinitionsToSteps(definitions config.Steps) (map[string]workflow.Step, error) {
	steps := map[string]workflow.Step{}

	for stepId, stepDefinition := range definitions {
		step, err := mapDefinitionToStep(stepId, stepDefinition)
		if err != nil {
			return nil, err
		}

		steps[stepId] = step
	}

	return steps, nil
}

func buildWorkflow(parsedConfig *config.Flow) (workflow.Workflow, error) {

	parsedSteps, err := mapStepsDefinitionsToSteps(parsedConfig.Workflow.Steps)
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
