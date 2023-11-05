package selflow

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/selflow/selflow/libs/core/sflog"
	"github.com/selflow/selflow/libs/core/workflow"
	"github.com/selflow/selflow/libs/selflow-daemon/config"
	"io"
	"log/slog"
)

const TerminationLogText = "===EOF==="

var TerminationLogBytes = []byte(fmt.Sprintf("%s\n", TerminationLogText))

type Selflow interface {
	StartRun(ctx context.Context, config *config.Flow) (runId string, err error)
}

type RunPersistence interface {
	SetRunState(runId string, state map[string]workflow.Status) error
	SetDependenciesState(runId string, dependencies map[workflow.Step][]workflow.Step) error
	SetRunStepDefinitions(runId string, stepDefinitions map[string]config.StepDefinition) error

	GetRunState(runId string) (map[string]workflow.Status, error)
	GetRunDependencies(runId string) (map[string][]string, error)
	GetRunStepDefinitions(runId string) (map[string][]byte, error)
}

type LogFactory interface {
	GetRunLogger(runId string) (io.Reader, io.WriteCloser, error)
	GetRunReader(runId string) (chan string, error)
}

type selflow struct {
	workflowBuilder WorkflowBuilder
	logFactory      LogFactory
	runPersistence  RunPersistence
}

func NewSelflow(workflowBuilder WorkflowBuilder, logFactory LogFactory, runPersistence RunPersistence) Selflow {
	return &selflow{
		workflowBuilder: workflowBuilder,
		logFactory:      logFactory,
		runPersistence:  runPersistence,
	}
}

func (s *selflow) StartRun(ctx context.Context, flow *config.Flow) (string, error) {
	runId := uuid.New().String()

	slog.DebugContext(ctx, "Building workflow")

	_, w, err := s.logFactory.GetRunLogger(runId)
	if err != nil {
		slog.ErrorContext(ctx, "Fail to get run logger", "error", err)
		return "", err
	}

	logger := sflog.NewLoggerWithWriter(w)
	ctx = sflog.ResetContextLogHandler(ctx, logger.Handler(), slog.String("workflowRun", runId))

	wf, err := s.workflowBuilder.BuildWorkflow(flow)
	if err != nil {
		slog.ErrorContext(ctx, "Fail to build workflow", "error", err)
		return "", err
	}

	simpleWf := wf.(*workflow.SimpleWorkflow)

	err = s.runPersistence.SetRunStepDefinitions(runId, flow.Workflow.Steps)
	if err != nil {
		return "", err
	}

	err = s.runPersistence.SetDependenciesState(runId, simpleWf.Dependencies)
	if err != nil {
		return "", err
	}

	logger.InfoContext(ctx, "here")
	slog.DebugContext(ctx, "Start workflow execution")
	slog.InfoContext(ctx, "Start workflow execution")

	go func(wf workflow.Workflow) {
		slog.InfoContext(ctx, "Start workflow execution")
		_, err := wf.Execute(ctx)
		if err != nil {
			slog.ErrorContext(ctx, "workflow execution failed", "error", err)
		} else {
			slog.InfoContext(ctx, "workflow execution succeeded")
		}
		if _, err = w.Write(TerminationLogBytes); err != nil {
			slog.ErrorContext(ctx, "fail to write closing log", "error", err)
		}
		if err = w.Close(); err != nil {
			slog.ErrorContext(ctx, "fail to close logger", "error", err)
		}
	}(wf)

	go func(simpleWf *workflow.SimpleWorkflow) {
		for state := range simpleWf.StateCh {
			err := s.runPersistence.SetRunState(runId, state)
			if err != nil {
				slog.WarnContext(ctx, "fail to save run state", "error", err)
			}
		}
	}(simpleWf)

	return runId, nil
}

var _ Selflow = &selflow{}
