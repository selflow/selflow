package docker_step

import (
	"context"
	"github.com/mitchellh/mapstructure"
	"github.com/selflow/selflow/internal/config"
	selflowRunnerProto "github.com/selflow/selflow/internal/selflow-runner-proto"
	"github.com/selflow/selflow/pkg/workflow"
)

type DockerStepConfig struct {
	Image        string            `mapstructure:"image"`
	Commands     string            `mapstructure:"commands"`
	Environments map[string]string `mapstructure:"environments"`
}

type DockerStep struct {
	workflow.SimpleStep
	DockerStepConfig
	ContainerSpawner selflowRunnerProto.ContainerSpawner
}

func (d *DockerStep) Execute(ctx context.Context) (map[string]string, error) {
	err := d.ContainerSpawner.SpawnContainer(ctx, "", d.Environments, d.Commands, d.Image)
	if err != nil {
		return nil, err
	}

	_, _ = d.SimpleStep.Execute(ctx)

	return map[string]string{}, nil
}

var _ workflow.Step = &DockerStep{}

func NewDockerStep(id string, definition config.StepDefinition, spawner selflowRunnerProto.ContainerSpawner) (*DockerStep, error) {
	dockerStepConfig := DockerStepConfig{}

	err := mapstructure.Decode(definition.With, &dockerStepConfig)
	if err != nil {
		return nil, err
	}

	return &DockerStep{
		SimpleStep: workflow.SimpleStep{
			Id:     id,
			Status: workflow.CREATED,
		},
		DockerStepConfig: dockerStepConfig,
		ContainerSpawner: spawner,
	}, nil
}
