package container

import (
	"github.com/mitchellh/mapstructure"
	workflow2 "github.com/selflow/selflow/libs/core/workflow"
	"github.com/selflow/selflow/libs/selflow-daemon/config"
	"github.com/selflow/selflow/libs/selflow-daemon/steps/conditional"
)

type StepMapper struct {
	ContainerSpawner ContainerSpawner
}

type DockerStepConfig struct {
	Image       string
	Commands    string
	Persistence map[string]string
}

func (c *StepMapper) mapStep(stepId string, definition config.StepDefinition) (workflow2.Step, error) {
	dockerStepConfig := DockerStepConfig{}

	delete(definition.With, "environment")

	err := mapstructure.Decode(definition.With, &dockerStepConfig)
	if err != nil {
		return nil, err
	}

	return &Step{
		containerSpawner: c.ContainerSpawner,
		SimpleStep:       workflow2.SimpleStep{Id: stepId, Status: workflow2.CREATED},
		config:           &dockerStepConfig,
	}, nil
}

func (c *StepMapper) MapStep(stepId string, definition config.StepDefinition) (workflow2.Step, error) {
	step, err := c.mapStep(stepId, definition)
	if err != nil {
		return nil, err
	}

	if definition.If != "" {
		step = conditional.NewConditionalStep(step, definition.If)
	}

	return step, nil
}
