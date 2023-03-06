package main

import (
	"context"
	"github.com/selflow/selflow/internal/config"
	selflowRunnerProto "github.com/selflow/selflow/internal/selflow-runner-proto"
	dockerStep "github.com/selflow/selflow/pkg/docker-step"
	"github.com/selflow/selflow/pkg/workflow"
	"log"
	"os"
)

type SelflowRunnerPlugin struct {
	containerSpawner selflowRunnerProto.ContainerSpawner
}

func (s *SelflowRunnerPlugin) InitRunner(ctx context.Context, spawner selflowRunnerProto.ContainerSpawner) error {
	s.containerSpawner = spawner

	configContent, err := os.ReadFile("/etc/selflow/config.json")
	if err != nil {
		return err
	}

	parsedConfig, err := config.Parse(configContent)
	if err != nil {
		return err
	}

	wf := workflow.MakeSimpleWorkflow(uint(len(parsedConfig.Workflow.Steps)))

	for i, stepDefinition := range parsedConfig.Workflow.Steps {

		step, err := dockerStep.NewDockerStep(i, stepDefinition, spawner)
		if err != nil {
			return err
		}

		err = wf.AddStep(step, []workflow.Step{})
		if err != nil {
			return err
		}
	}

	err = wf.Init(ctx)

	if err != nil {
		return err
	}

	workflowExecutionResults, err := wf.Execute(ctx)

	if err != nil {
		return err
	}
	log.Println(workflowExecutionResults)

	log.Println(parsedConfig.Metadata)
	return nil
}
