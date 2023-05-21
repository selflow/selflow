package container

import (
	"github.com/mitchellh/mapstructure"
	"github.com/selflow/selflow/internal/config"
	"github.com/selflow/selflow/pkg/workflow"
)

type StepMapper struct {
	ContainerSpawner ContainerSpawner
}

func (c *StepMapper) MapStep(stepId string, definition config.StepDefinition) (workflow.Step, error) {
	dockerStepConfig := ContainerConfig{}

	delete(definition.With, "environment")

	err := mapstructure.Decode(definition.With, &dockerStepConfig)
	if err != nil {
		return nil, err
	}

	return &Step{
		containerSpawner: c.ContainerSpawner,
		SimpleStep:       workflow.SimpleStep{Id: stepId, Status: workflow.CREATED},
		config:           &dockerStepConfig,
	}, nil
}
