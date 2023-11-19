package models

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/selflow/selflow/libs/selflow-daemon/sfenvironment"
	"io"
	"strings"
)

type WorkflowLogMessage struct {
	Message string `json:"msg"`
	StepId  string `json:"StepId"`
	Channel string `json:"channel"`
	Status  string `json:"stepStatus"`
	Level   string `json:"level"`
	error   error
}

type logModel struct {
	ctx context.Context

	logs       []*WorkflowLogMessage
	maxLogSize int
	scanner    *bufio.Scanner
}

func NewLogModel(ctx context.Context, size int, reader io.Reader) tea.Model {
	return logModel{
		ctx:        ctx,
		logs:       make([]*WorkflowLogMessage, size),
		maxLogSize: size,
		scanner:    bufio.NewScanner(reader),
	}
}

func (m logModel) startLogger() tea.Msg {
	if m.scanner.Scan() {
		if sfenvironment.UseJsonLogs {
			fmt.Println(m.scanner.Text())
		}
		msg := WorkflowLogMessage{}
		if err := json.Unmarshal(m.scanner.Bytes(), &msg); err != nil {
			return WorkflowLogMessage{error: err}
		}

		return msg
	}
	return WorkflowLogMessage{error: m.scanner.Err()}
}

func (m logModel) Init() tea.Cmd {
	return m.startLogger
}

func (m logModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case WorkflowLogMessage:
		messagesTexts := strings.Split(msg.Message, "\n")
		msgs := make([]*WorkflowLogMessage, 0, m.maxLogSize)

		for _, logLine := range messagesTexts[:min(len(messagesTexts), m.maxLogSize)] {
			if logLine == "" {
				continue
			}
			msgs = append(msgs, &WorkflowLogMessage{
				Message: logLine,
				StepId:  msg.StepId,
				Channel: msg.Channel,
			})
		}

		if msg.Message != "" {
			m.logs = append(m.logs[len(msgs):], msgs...)
		}
		return m, m.startLogger
	}
	return m, nil
}

func (m logModel) View() string {
	logViewBuffer := ""

	for _, l := range m.logs {
		if l == nil {
			logViewBuffer += "........................\n"
		} else {
			logViewBuffer += fmt.Sprintf("%s |\t %s\n", l.StepId, l.Message)
		}
	}

	return logViewBuffer
}
