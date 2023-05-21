package container

import (
	"context"
	"errors"
	"github.com/hashicorp/go-hclog"
	"github.com/selflow/selflow/pkg/sflog"
	"github.com/selflow/selflow/pkg/workflow"
)

var ContainerExitedNon0StatusCodeError = errors.New("container exited with a non-zero status code")

type Step struct {
	workflow.SimpleStep
	containerSpawner ContainerSpawner
	config           *ContainerConfig
	output           map[string]string
}

func (step *Step) GetOutput() map[string]string {
	return step.output
}

func (step *Step) Execute(ctx context.Context) (map[string]string, error) {
	step.SetStatus(workflow.RUNNING)

	logger := sflog.LoggerFromContext(ctx)
	stepLogger := logger.Named(step.GetId())

	needs := ctx.Value(workflow.StepOutputContextKey).(map[string]map[string]string)

	var err error
	var image string
	var commands string

	if image, err = withTemplate(step.config.Image, needs); err != nil {
		return nil, err
	}

	if commands, err = withTemplate(step.config.Commands, needs); err != nil {
		return nil, err
	}

	containerId, err := step.containerSpawner.StartContainerDetached(ctx, &ContainerConfig{
		Image:               image,
		Commands:            commands,
		ContainerName:       step.config.ContainerName,
		ContainerLogsWriter: step.config.ContainerLogsWriter,
		Environment:         step.config.Environment,
		Entrypoint:          step.config.Entrypoint,
		Mounts:              step.config.Mounts,
		OpenPorts:           step.config.OpenPorts,
		Networks:            step.config.Networks,
	})
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
