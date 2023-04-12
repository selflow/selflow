package main

import (
	"context"
	"github.com/mitchellh/mapstructure"
	"github.com/selflow/selflow/internal/config"
	"github.com/selflow/selflow/pkg/workflow"
	"log"
	"os"
)

type ContainerStep struct {
	workflow.SimpleStep
	containerSpawner ContainerSpawner
	config           *ContainerConfig
}

type ContainerStepMapper struct {
	containerSpawner ContainerSpawner
}

func (c *ContainerStepMapper) MapStep(stepId string, definition config.StepDefinition) (workflow.Step, error) {
	dockerStepConfig := ContainerConfig{}

	delete(definition.With, "environment") // TODO : handle environment

	err := mapstructure.Decode(definition.With, &dockerStepConfig)
	if err != nil {
		return nil, err
	}

	return &ContainerStep{
		containerSpawner: c.containerSpawner,
		SimpleStep:       workflow.SimpleStep{Id: stepId, Status: workflow.CREATED},
		config:           &dockerStepConfig,
	}, nil
}

func (step *ContainerStep) Execute(ctx context.Context) (map[string]string, error) {
	step.SetStatus(workflow.RUNNING)

	containerId, err := step.containerSpawner.StartContainerDetached(ctx, step.config)
	if err != nil {
		return nil, err
	}

	go func() {
		err := step.containerSpawner.TransferContainerLogs(ctx, containerId, os.Stdout)
		if err != nil {

			log.Printf("[WARN] fail to transfer container logs : %v", err)
		}
	}()

	exitCode, err := step.containerSpawner.WaitContainer(ctx, containerId)
	if err != nil {
		return nil, err
	}

	if exitCode != 0 {
		log.Printf("[INFO] fail to transfer container logs : %v", err)
		return nil, err
	}

	return map[string]string{}, nil
}
