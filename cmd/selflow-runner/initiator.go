package main

import (
	"context"
	"github.com/selflow/selflow/internal/config"
	selflowRunnerProto "github.com/selflow/selflow/internal/selflow-runner-proto"
	"github.com/selflow/selflow/pkg/workflow"
	"log"
	"os"
)

type SelflowRunnerPlugin struct {
	containerSpawner selflowRunnerProto.ContainerSpawner
	workflowBuilder  interface {
		BuildWorkflow(parsedConfig *config.Flow) (workflow.Workflow, error)
	}
	configFileLocation string
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

	log.Println("[DEBUG] Parsing configuration")

	configContent, err := os.ReadFile(s.configFileLocation)
	if err != nil {
		return err
	}

	parsedConfig, err := config.Parse(configContent)
	if err != nil {
		return err
	}

	wf, err := s.workflowBuilder.BuildWorkflow(parsedConfig)
	if err != nil {
		return err
	}

	err = wf.Init()

	if err != nil {
		return err
	}

	ctxWithSpawner := context.WithValue(ctx, selflowRunnerProto.ContainerSpawnerContextKey, containerSpawner)

	log.Println("[DEBUG] Start workflow execution")

	_, err = wf.Execute(ctxWithSpawner)

	if err != nil {
		return err
	}

	return nil
}
