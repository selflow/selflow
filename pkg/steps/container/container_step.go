package container

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/selflow/selflow/pkg/sflog"
	"github.com/selflow/selflow/pkg/workflow"
)

var ContainerExitedNon0StatusCodeError = errors.New("container exited with a non-zero status code")

type Step struct {
	workflow.SimpleStep
	containerSpawner ContainerSpawner
	config           *DockerStepConfig
	output           map[string]string
}

func (step *Step) GetOutput() map[string]string {
	return step.output
}

func (step *Step) Execute(ctx context.Context) (map[string]string, error) {
	step.SetStatus(workflow.RUNNING)

	runId := ctx.Value(workflow.RunIdContextKey{}).(string)
	logger := sflog.LoggerFromContext(ctx)
	stepLogger := logger.Named(step.GetId())

	needs := ctx.Value(workflow.StepOutputContextKey).(map[string]map[string]string)

	var err error

	containerConfig := &ContainerConfig{}

	if containerConfig.Image, err = withTemplate(step.config.Image, needs); err != nil {
		return nil, err
	}

	if containerConfig.Commands, err = withTemplate(step.config.Commands, needs); err != nil {
		return nil, err
	}

	if step.config.Persistence != nil {
		mounts := make([]Mountable, 0, len(step.config.Persistence))

		for artifactName, destDir := range step.config.Persistence {
			mounts = append(mounts, Mount{
				ArtifactName: fmt.Sprintf("sf-%s-%s", runId, artifactName),
				Destination:  destDir,
			})
		}

		containerConfig.Mounts = mounts
	}

	containerId, err := step.containerSpawner.StartContainerDetached(ctx, containerConfig)
	if err != nil {
		return nil, err
	}

	logWriter := &writerWithOutput{
		Writer: stepLogger.StandardWriter(&hclog.StandardLoggerOptions{
			ForceLevel: hclog.Debug,
		}),
		output: map[string]string{},
	}

	err = step.containerSpawner.TransferContainerLogs(ctx, containerId, logWriter)
	if err != nil {

		logger.Warn("fail to transfer container logs", "error", err)
	}

	exitCode, err := step.containerSpawner.WaitContainer(ctx, containerId)
	if err != nil {
		return nil, err
	}

	if exitCode != 0 {
		stepLogger.Error("container exited with status", "ExitCode", exitCode)
		return nil, ContainerExitedNon0StatusCodeError
	}
	step.output = logWriter.output

	return logWriter.output, nil
}
