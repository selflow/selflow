package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
	"github.com/selflow/selflow/libs/core/selflow"
	"github.com/selflow/selflow/libs/core/workflow"
	"github.com/selflow/selflow/libs/selflow-cli/steps/localexec"
	"github.com/selflow/selflow/libs/selflow-daemon/config"
	"github.com/selflow/selflow/libs/selflow-daemon/container-spawner/docker"
	"github.com/selflow/selflow/libs/selflow-daemon/steps/container"
	"github.com/spf13/cobra"
	"io"
	"log"
	"log/slog"
	"os"
)

func NewExecCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "exec",
		Short: "Exec a workflow without the daemon",
		Long: "Exec a workflow without the daemon. " +
			"Only use it for testing or local purpose. " +
			"Not fit in production." +
			"Be carefully to the commands that will be triggered",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := execWorkflowLocally(args[0]); err != nil {
				if _, err = fmt.Fprintf(os.Stderr, "%v\n", err); err != nil {
					panic(err)
				}
				os.Exit(1)
			}
		},
	}
}

func execWorkflowLocally(configFile string) error {
	ctx := context.Background()

	configContent, err := os.ReadFile(configFile)
	if err != nil {
		slog.ErrorContext(ctx, "Fail to read file", "error", err)
		return err
	}

	flow, err := config.Parse(configContent)
	if err != nil {
		slog.ErrorContext(ctx, "Fail to parse configuration", "error", err)
		return err
	}

	slog.DebugContext(ctx, "Connecting to Docker daemon")
	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		slog.ErrorContext(ctx, "Fail to connect to Docker Daemon")
		return err
	}

	workflowBuilder := selflow.WorkflowBuilder{
		StepMappers: []selflow.StepMapper{
			&container.StepMapper{
				ContainerSpawner: docker.NewSpawner(dockerClient),
			},
			&localexec.StepMapper{},
		},
	}

	self := selflow.NewSelflow(workflowBuilder, &StdLogFactory{}, BlankPersistence{})

	runId, terminatedWorkflow, err := self.StartRun(ctx, flow)
	if err != nil {
		slog.ErrorContext(ctx, "Error Starting the workflow")
		return err
	}

	workflowSucceeded := <-terminatedWorkflow
	if workflowSucceeded {
		log.Printf("Success !\n")
	} else {
		log.Printf("Fail\n")
	}

	slog.InfoContext(ctx, "Workflow Started", "runId", runId)
	return nil
}

type BlankPersistence struct {
}

func (b BlankPersistence) SetRunState(_ string, _ map[string]workflow.Status) error {
	return nil
}

func (b BlankPersistence) SetDependenciesState(_ string, _ map[workflow.Step][]workflow.Step) error {
	return nil
}

func (b BlankPersistence) SetRunStepDefinitions(_ string, _ map[string]config.StepDefinition) error {
	return nil
}

func (b BlankPersistence) GetRunState(_ string) (map[string]workflow.Status, error) {
	return map[string]workflow.Status{}, nil
}

func (b BlankPersistence) GetRunDependencies(_ string) (map[string][]string, error) {
	return map[string][]string{}, nil
}

func (b BlankPersistence) GetRunStepDefinitions(_ string) (map[string][]byte, error) {
	return map[string][]byte{}, nil
}

type StdLogFactory struct {
}

func (s StdLogFactory) GetRunLogger(runId string) (io.Reader, io.WriteCloser, error) {
	return os.Stdout, os.Stdout, nil
}

func (s StdLogFactory) GetRunReader(runId string) (chan string, error) {
	return nil, nil
}
