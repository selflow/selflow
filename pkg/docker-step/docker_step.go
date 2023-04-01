package docker_step

import (
	"context"
	"errors"
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
}

var MissingContainerSpawnerContextKey = errors.New("container spawner context key is missing")

func (d *DockerStep) Execute(ctx context.Context) (map[string]string, error) {

	containerSpawner, ok := ctx.Value(selflowRunnerProto.ContainerSpawnerContextKey).(selflowRunnerProto.ContainerSpawner)
	if !ok {
		return nil, MissingContainerSpawnerContextKey
	}

	err := containerSpawner.SpawnContainer(ctx, d.Id, d.Environments, d.Commands, d.Image)
	if err != nil {
		return nil, err
	}

	_, _ = d.SimpleStep.Execute(ctx)

	return map[string]string{}, nil
}

var _ workflow.Step = &DockerStep{}

func NewDockerStep(id string, definition config.StepDefinition) (workflow.Step, error) {
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
	}, nil
}
