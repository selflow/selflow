package models

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/indent"
	"github.com/selflow/selflow/libs/core/workflow"
	"github.com/selflow/selflow/libs/selflow-daemon/config"
	"io"
)

// RunModel is a bubbletea model that follows a selflow run and shows the running steps and the associated logs
type RunModel struct {
	// display purpose
	spinner         spinner.Model
	quitting        bool
	endMsg          processFinishedMsg
	logModel        tea.Model
	stepStatusModel StepStatusModel

	// Workflow monitoring
	workflowTerminated chan bool
}

func NewRunModel(ctx context.Context, workflowTerminated chan bool, workflowLogs io.Reader, stepStatusCh chan StepStatus) *RunModel {
	const showLastResults = 15

	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("206"))

	return &RunModel{
		spinner:            sp,
		logModel:           NewLogModel(ctx, showLastResults, workflowLogs),
		stepStatusModel:    NewStepStatusModel(ctx, sp, stepStatusCh),
		workflowTerminated: workflowTerminated,
	}
}

type processFinishedMsg string

// runWorkflow is a bubbletea command that executes a selflow workflow and returns a message when it ends with its status.
// the resulting message can then be used to update the bubbletea interface
func (m *RunModel) runWorkflow() tea.Msg {
	workflowTerminatedSuccessfully := <-m.workflowTerminated
	if workflowTerminatedSuccessfully {
		return processFinishedMsg("Success")
	} else {
		return processFinishedMsg("Fail")
	}
}

// Init starts all bubbletea commands that needs to run in the "background" of the execution
func (m *RunModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		m.runWorkflow,
		m.logModel.Init(),
		m.stepStatusModel.Init(),
	)
}

// Update is called everytime a bubbletea command ends. It handles the different events by updating the models.
func (m *RunModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// A key is pressed, stop the process
		m.quitting = true
		m.endMsg = "cancelled"
		return m, tea.Quit

	case processFinishedMsg:
		// The workflow has ended, stop the process
		m.quitting = true
		m.endMsg = msg
		return m, tea.Quit

		// Spinner
	case spinner.TickMsg:
		// the spinner frame needs to change
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		m.stepStatusModel.Spinner = m.spinner
		return m, cmd

		// Logs
	case WorkflowLogMessage:
		// received new log, forward it to the log model
		var cmd tea.Cmd
		m.logModel, cmd = m.logModel.Update(msg)
		return m, cmd

		// Step Status
	case StepStatus:
		// the status of a step changed, update the StepStatusModel model
		var cmd tea.Cmd
		m.stepStatusModel, cmd = m.stepStatusModel.Update(msg)
		return m, cmd

	default:
		// uncaught event, do nothing
		return m, nil
	}
}

// View creates a string describing what should be rendered in the terminal.
func (m *RunModel) View() string {
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

// LivePersistence implements the selflow persistence interface to handle changes in the steps.
// Each step change is added in a channel to be handled in the log_model
type LivePersistence struct {
	StepStatusCh chan StepStatus
}

func (b LivePersistence) SetRunState(_ string, steps map[string]workflow.Status) error {
	for stepId, status := range steps {
		b.StepStatusCh <- StepStatus{StepId: stepId, Status: status.GetName()}
	}
	return nil
}

func (b LivePersistence) SetDependenciesState(_ string, steps map[workflow.Step][]workflow.Step) error {
	for closingStep := range steps {
		b.StepStatusCh <- StepStatus{StepId: closingStep.GetId(), Status: closingStep.GetStatus().GetName()}
	}
	return nil
}

func (b LivePersistence) SetRunStepDefinitions(_ string, _ map[string]config.StepDefinition) error {
	return nil
}
