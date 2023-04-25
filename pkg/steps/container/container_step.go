package container

import (
	"context"
	"errors"
	"github.com/hashicorp/go-hclog"
	"github.com/mitchellh/mapstructure"
	"github.com/selflow/selflow/internal/config"
	"github.com/selflow/selflow/pkg/workflow"
	"log"
)

var ContainerExitedNon0StatusCodeError = errors.New("container exited with a non-zero status code")

type Step struct {
	workflow.SimpleStep
	containerSpawner ContainerSpawner
	config           *ContainerConfig
}

type StepMapper struct {
	ContainerSpawner ContainerSpawner
}

func (c *StepMapper) MapStep(stepId string, definition config.StepDefinition) (workflow.Step, error) {
	dockerStepConfig := ContainerConfig{}

	delete(definition.With, "environment") // TODO : handle environment

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

func (step *Step) Execute(ctx context.Context) (map[string]string, error) {
	step.SetStatus(workflow.RUNNING)

	stepLogger := hclog.L().Named(step.GetId())

	containerId, err := step.containerSpawner.StartContainerDetached(ctx, step.config)
	if err != nil {
		return nil, err
	}

	err = step.containerSpawner.TransferContainerLogs(ctx, containerId, stepLogger.StandardWriter(&hclog.StandardLoggerOptions{
		ForceLevel: hclog.Debug,
	}))
	if err != nil {

		log.Printf("[WARN] fail to transfer container logs : %v", err)
	}

	exitCode, err := step.containerSpawner.WaitContainer(ctx, containerId)
	if err != nil {
		return nil, err
	}

	if exitCode != 0 {
		stepLogger.Error("container exited with status", "ExitCode", exitCode)
		return nil, ContainerExitedNon0StatusCodeError
	}

	return map[string]string{}, nil
}
