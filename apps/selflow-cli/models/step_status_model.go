package models

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type StepStatus struct {
	StepId string
	Status string
}

type StepStatusModel struct {
	Spinner spinner.Model
	ctx     context.Context

	stepStatusChanged chan StepStatus
	stepStatus        []StepStatus
}

func NewStepStatusModel(ctx context.Context, spinner spinner.Model, stepStatusChanged chan StepStatus) StepStatusModel {
	return StepStatusModel{
		Spinner:           spinner,
		ctx:               ctx,
		stepStatusChanged: stepStatusChanged,
	}
}

func (m StepStatusModel) listenForStepStatus() tea.Msg {
	step, isOpen := <-m.stepStatusChanged
	if isOpen {
		return step
	}
	return nil
}

func (m StepStatusModel) Init() tea.Cmd {
	return m.listenForStepStatus
}

func (m StepStatusModel) Update(msg tea.Msg) (StepStatusModel, tea.Cmd) {
	switch msg := msg.(type) {
	case StepStatus:
		a := m.stepStatus

		inLoop := false
		for i, s := range a {
			if s.StepId == msg.StepId {
				a[i] = msg
				inLoop = true
				break
			}
		}

		if !inLoop {
			a = append(a, msg)
		}

		m.stepStatus = a
		return m, m.listenForStepStatus

	}
	return m, nil
}

func (m StepStatusModel) View() string {
	stepStateViewBuffer := ""
	for _, s := range m.stepStatus {

		prefix := "?"

		switch strings.ToLower(s.Status) {
		case "success":
			prefix = "✅"
			break
		case "error":
			prefix = "❌"
			break
		case "cancelled":
			prefix = "🚫"
			break
		default:
			prefix = m.Spinner.View()
		}

		stepStateViewBuffer += fmt.Sprintf("%2s %s\n", prefix, s.StepId)
	}

	return stepStateViewBuffer
}
