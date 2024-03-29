package localexec

import (
	"errors"
	"github.com/mitchellh/mapstructure"
	"github.com/selflow/selflow/libs/core/workflow"
	"github.com/selflow/selflow/libs/selflow-daemon/config"
	"github.com/selflow/selflow/libs/selflow-daemon/steps/conditional"
)

type StepMapper struct {
}

func (c *StepMapper) mapStep(stepId string, definition config.StepDefinition) (workflow.Step, error) {
	dockerStepConfig := Config{}

	delete(definition.With, "environment")

	err := mapstructure.Decode(definition.With, &dockerStepConfig)
	if err != nil {
		return nil, err
	}

	return &Step{
		SimpleStep: workflow.SimpleStep{Id: stepId, Status: workflow.PENDING},
		config:     dockerStepConfig,
	}, nil
}

func (c *StepMapper) MapStep(stepId string, definition config.StepDefinition) (workflow.Step, error) {
	if definition.Kind != "localexec" {
		return nil, errors.New("invalid kind")
	}

	step, err := c.mapStep(stepId, definition)
	if err != nil {
		return nil, err
	}

	if definition.If != "" {
		step = conditional.NewConditionalStep(step, definition.If)
	}

	return step, nil
}
