package main

import (
	"context"
	"github.com/docker/docker/client"
	"github.com/hashicorp/go-hclog"
	"github.com/selflow/selflow/internal/config"
	"log"
	"os"
)

func setupLogger() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:            "selflow-daemon",
		Output:          nil,
		JSONFormat:      false,
		IncludeLocation: false,
		TimeFormat:      "2006-01-02 15:04:05",
		Color:           hclog.ForceColor,
		Level:           hclog.Debug,
	})

	hclog.SetDefault(logger)

	log.SetOutput(logger.StandardWriter(&hclog.StandardLoggerOptions{InferLevels: true}))
	log.SetPrefix("")
	log.SetFlags(0)
}

func main() {
	setupLogger()

	configAsBytes, err := os.ReadFile("/home/anthony/Projets/selflow/internal/config/testdata/flow.yaml")
	if err != nil {
		panic(err)
	}

	flow, err := config.Parse(configAsBytes)
	if err != nil {
		panic(err)
	}

	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	workflowBuilder := WorkflowBuilder{
		stepMappers: []StepMapper{
			&ContainerStepMapper{
				containerSpawner: &dockerSpawner{
					dockerClient,
				},
			},
		},
	}

	wf, err := workflowBuilder.BuildWorkflow(flow)
	if err != nil {
		panic(err)
	}
	_, err = wf.Execute(context.Background())
	if err != nil {
		panic(err)
	}

}
