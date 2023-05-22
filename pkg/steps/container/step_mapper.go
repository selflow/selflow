package container

import (
	"github.com/mitchellh/mapstructure"
	"github.com/selflow/selflow/internal/config"
	"github.com/selflow/selflow/pkg/steps/conditional"
	"github.com/selflow/selflow/pkg/workflow"
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
		SimpleStep:       workflow.SimpleStep{Id: stepId, Status: workflow.CREATED},
		config:           &dockerStepConfig,
	}, nil
}

func (c *StepMapper) MapStep(stepId string, definition config.StepDefinition) (workflow.Step, error) {
	step, err := c.mapStep(stepId, definition)
	if err != nil {
		return nil, err
	}

	if definition.If != "" {
		step = conditional.NewConditionalStep(step, definition.If)
	}

	return step, nil
}
