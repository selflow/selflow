package main

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/docker/client"
	"github.com/muesli/reflow/indent"
	"github.com/selflow/selflow/apps/selflow-cli/models"
	"github.com/selflow/selflow/libs/core/selflow"
	"github.com/selflow/selflow/libs/core/sflog"
	"github.com/selflow/selflow/libs/core/workflow"
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

type runModel struct {
	spinner         spinner.Model
	selflow         selflow.Selflow
	flow            *config.Flow
	writer          io.Writer
	quitting        bool
	ctx             context.Context
	logModel        tea.Model
	stepStatusModel models.StepStatusModel
	endMsg          processFinishedMsg
}

func newRunModel(ctx context.Context, workflowBuilder selflow.WorkflowBuilder, flow *config.Flow) *runModel {

	reader, writer := io.Pipe()
	const showLastResults = 15

	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("206"))

	stepStatusCh := make(chan models.StepStatus)

	return &runModel{
		spinner:         sp,
		writer:          writer,
		flow:            flow,
		selflow:         selflow.NewSelflow(workflowBuilder, LivePersistence{stepStatusCh}),
		ctx:             ctx,
		logModel:        models.NewLogModel(ctx, showLastResults, reader),
		stepStatusModel: models.NewStepStatusModel(ctx, sp, stepStatusCh),
	}
}

type processFinishedMsg string

func (m *runModel) startWorkflow() tea.Msg {
	if sfenvironment.UseJsonLogs {
		slog.InfoContext(m.ctx, "Start Workflow")
	}
	runCtx := sflog.ResetContextLogHandler(m.ctx, slog.NewJSONHandler(m.writer, &slog.HandlerOptions{Level: slog.LevelDebug}))

	_, err := m.selflow.StartRun(runCtx, m.flow)

	if sfenvironment.UseJsonLogs {
		if err == nil {
			slog.InfoContext(m.ctx, "Workflow executed successfully")
		} else {
			slog.InfoContext(m.ctx, "Workflow terminated with an error", "error", err)
		}
	}

	if err == nil {
		return processFinishedMsg("Success")
	} else {
		return processFinishedMsg("Fail")
	}
}

func (m *runModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		m.startWorkflow,
		m.logModel.Init(),
		m.stepStatusModel.Init(),
	)
}

func (m *runModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.quitting = true
		m.endMsg = "cancelled"
		return m, tea.Quit

	case processFinishedMsg:
		m.quitting = true
		m.endMsg = msg
		return m, tea.Quit

		// Spinner
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		m.stepStatusModel.Spinner = m.spinner
		return m, cmd

		// Logs
	case models.WorkflowLogMessage:
		var cmd tea.Cmd
		m.logModel, cmd = m.logModel.Update(msg)
		return m, cmd

		// Step Status
	case models.StepStatus:
		var cmd tea.Cmd
		m.stepStatusModel, cmd = m.stepStatusModel.Update(msg)
		return m, cmd

	default:
		return m, nil
	}
}

func (m *runModel) View() string {
	buffer := "\n"

	if !m.quitting {
		buffer += m.spinner.View() + " Running Workflow...\n\n"
	}

	stepStateViewBuffer := m.stepStatusModel.View()

	logViewBuffer := m.logModel.View()

	buffer += lipgloss.JoinHorizontal(lipgloss.Top, stepStateViewBuffer, indent.String(logViewBuffer, 4))

	if m.quitting {
		buffer += fmt.Sprintf("\n\n Workflow terminated with status [%s]\n\n", m.endMsg)
	}

	return buffer
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

	model := newRunModel(ctx, workflowBuilder, flow)

	var bubbleteaOptions []tea.ProgramOption
	// Handle sessions without tty
	if !term.IsTerminal(int(os.Stdout.Fd())) {
		slog.DebugContext(ctx, "Not tty detected")
		bubbleteaOptions = append(bubbleteaOptions, tea.WithInput(nil))
	}
	if sfenvironment.UseJsonLogs {
		bubbleteaOptions = append(bubbleteaOptions, tea.WithoutRenderer())
	}

	p := tea.NewProgram(model, bubbleteaOptions...)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting Bubble Tea program:", err)
		os.Exit(1)
	}
	return nil
}

type LivePersistence struct {
	stepStatusCh chan models.StepStatus
}

func (b LivePersistence) SetRunState(_ string, steps map[string]workflow.Status) error {
	for stepId, status := range steps {
		b.stepStatusCh <- models.StepStatus{StepId: stepId, Status: status.GetName()}
	}
	return nil
}

func (b LivePersistence) SetDependenciesState(_ string, steps map[workflow.Step][]workflow.Step) error {
	for closingStep := range steps {
		b.stepStatusCh <- models.StepStatus{StepId: closingStep.GetId(), Status: closingStep.GetStatus().GetName()}
	}
	return nil
}

func (b LivePersistence) SetRunStepDefinitions(_ string, _ map[string]config.StepDefinition) error {
	return nil
}
