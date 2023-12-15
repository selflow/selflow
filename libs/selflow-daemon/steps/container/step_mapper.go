package container

import (
	"errors"
	"github.com/mitchellh/mapstructure"
	"github.com/selflow/selflow/libs/core/workflow"
	"github.com/selflow/selflow/libs/selflow-daemon/config"
)

type StepMapper struct {
	ContainerSpawner ContainerSpawner
}

type DockerStepConfig struct {
	Image       string
	Commands    string
	Persistence map[string]string
}

func (c *StepMapper) mapStep(stepId string, definition config.StepDefinition) (workflow.Step, error) {
	dockerStepConfig := DockerStepConfig{}

	delete(definition.With, "environment")

	err := mapstructure.Decode(definition.With, &dockerStepConfig)
	if err != nil {
		return nil, err
	}

	return &Step{
		containerSpawner: c.ContainerSpawner,
		SimpleStep:       workflow.SimpleStep{Id: stepId, Status: workflow.PENDING},
		config:           &dockerStepConfig,
	}, nil
}

func (c *StepMapper) MapStep(stepId string, definition config.StepDefinition) (workflow.Step, error) {
	if definition.Kind != "docker" {
		return nil, errors.New("invalid kind")
	}

	step, err := c.mapStep(stepId, definition)
	if err != nil {
		return nil, err
	}

	return step, nil
}
