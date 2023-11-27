package main

import (
	"context"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/docker/docker/client"
	"github.com/selflow/selflow/apps/selflow-cli/models"
	"github.com/selflow/selflow/libs/core/selflow"
	"github.com/selflow/selflow/libs/core/sflog"
	"github.com/selflow/selflow/libs/selflow-cli/steps/localexec"
	"github.com/selflow/selflow/libs/selflow-daemon/config"
	"github.com/selflow/selflow/libs/selflow-daemon/container-spawner/docker"
	"github.com/selflow/selflow/libs/selflow-daemon/sfenvironment"
	"github.com/selflow/selflow/libs/selflow-daemon/steps/container"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"io"
	"log/slog"
	"os"
	"path"
)

func NewExecCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "exec [filename]",
		Short: "Execute a workflow locally without the daemon",
		Long: "Execute a workflow locally without the daemon.\n" +
			"It allows the execution of local command. Be careful what you are executing :)\n\n" +
			"Supported steps are:\n" +
			"- docker\n" +
			"- localexec",
		Args:       cobra.MinimumNArgs(1),
		ArgAliases: []string{"configFile"},
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

	//------------------------------------//
	//--- Parse the configuration file ---//
	//------------------------------------//
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

	//--------------------------//
	//--- Initialize plugins ---//
	//--------------------------//

	slog.DebugContext(ctx, "Connecting to Docker daemon")
	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		slog.ErrorContext(ctx, "Fail to connect to Docker Daemon")
		return err
	}

	tmpDirectory := path.Join(tmpFileDirectory, "tmp")

	workflowBuilder := selflow.WorkflowBuilder{
		StepMappers: []selflow.StepMapper{
			&container.StepMapper{
				ContainerSpawner: docker.NewSpawner(dockerClient, tmpDirectory, tmpDirectory),
			},
			&localexec.StepMapper{},
		},
	}

	//----------------------//
	//--- Start Workflow ---//
	//----------------------//

	workflowLogsReader, workflowLogsWriter := io.Pipe()
	workflowTerminated := make(chan bool, 1)
	stepStatusCh := make(chan models.StepStatus)

	go func(ctx context.Context) {
		ctx = sflog.ResetContextLogHandler(ctx, slog.NewJSONHandler(workflowLogsWriter, &slog.HandlerOptions{Level: slog.LevelDebug}))

		sf := selflow.NewSelflow(workflowBuilder, models.LivePersistence{StepStatusCh: stepStatusCh})

		_, err = sf.StartRun(ctx, flow)
		if sfenvironment.UseJsonLogs {
			if err == nil {
				slog.InfoContext(ctx, "Workflow executed successfully")
			} else {
				slog.ErrorContext(ctx, "Workflow terminated with an error", "error", err)
			}
		}
		workflowTerminated <- err == nil
	}(ctx)

	//----------------------------//
	//--- Initialize Bubbletea ---//
	//----------------------------//

	model := models.NewRunModel(ctx, workflowTerminated, workflowLogsReader, stepStatusCh)

	var bubbleteaOptions []tea.ProgramOption
	// Handle sessions without tty
	if !term.IsTerminal(int(os.Stdout.Fd())) {
		slog.DebugContext(ctx, "Not tty detected")
		bubbleteaOptions = append(bubbleteaOptions, tea.WithInput(nil))
	}
	if sfenvironment.UseJsonLogs {
		bubbleteaOptions = append(bubbleteaOptions, tea.WithoutRenderer())
	}

	//---------------------//
	//--- Start process ---//
	//---------------------//

	p := tea.NewProgram(model, bubbleteaOptions...)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting Bubble Tea program:", err)
		os.Exit(1)
	}
	return nil
}
