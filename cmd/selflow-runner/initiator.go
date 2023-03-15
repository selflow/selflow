package main

import (
	"context"
	"github.com/selflow/selflow/internal/config"
	selflowRunnerProto "github.com/selflow/selflow/internal/selflow-runner-proto"
	"log"
	"os"
)

type SelflowRunnerPlugin struct {
	containerSpawner selflowRunnerProto.ContainerSpawner
}

func (s *SelflowRunnerPlugin) InitRunner(ctx context.Context, containerSpawner selflowRunnerProto.ContainerSpawner) error {
	err := s.initRunner(ctx, containerSpawner)
	if err != nil {
		log.Printf("[ERROR] %v\n", err.Error())
	}
	return err
}

func (s *SelflowRunnerPlugin) initRunner(ctx context.Context, containerSpawner selflowRunnerProto.ContainerSpawner) error {
	s.containerSpawner = containerSpawner

	configContent, err := os.ReadFile("/etc/selflow/config.json")
	if err != nil {
		return err
	}

	parsedConfig, err := config.Parse(configContent)
	if err != nil {
		return err
	}

	wf, err := buildWorkflow(parsedConfig)
	if err != nil {
		return err
	}

	ctxWithSpawner := context.WithValue(ctx, selflowRunnerProto.ContainerSpawnerContextKey, containerSpawner)
	err = wf.Init(ctxWithSpawner)

	if err != nil {
		return err
	}

	workflowExecutionResults, err := wf.Execute(ctxWithSpawner)

	if err != nil {
		return err
	}
	log.Println(workflowExecutionResults)

	return nil
}
