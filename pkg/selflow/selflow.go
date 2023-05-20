package selflow

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/selflow/selflow/internal/config"
	"github.com/selflow/selflow/pkg/sflog"
	"github.com/selflow/selflow/pkg/workflow"
	"io"
)

const TerminationLogText = "===EOF==="

var TerminationLogBytes = []byte(fmt.Sprintf("%s\n", TerminationLogText))

type Selflow interface {
	StartRun(ctx context.Context, config *config.Flow) (runId string, err error)
}

type RunPersistence interface {
	SetRunState(runId string, state map[string]workflow.Status) error
	GetRunState(runId string) (map[string]workflow.Status, error)
	GetRunDependencies(runId string) (map[string][]string, error)
	SetDependenciesState(runId string, dependencies map[workflow.Step][]workflow.Step) error
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

	wf, err := s.workflowBuilder.BuildWorkflow(flow)
	if err != nil {
		return "", err
	}

	ctxLogger := sflog.LoggerFromContext(ctx)

	_, w, err := s.logFactory.GetRunLogger(runId)
	if err != nil {
		ctxLogger.Error("fail to create logger", "error", err)
		return "", err
	}

	logger := sflog.LoggerWithWriter(ctxLogger.Name(), w).Named(runId)

	runCtx := sflog.ContextWithLogger(context.Background(), logger)

	simpleWf := wf.(*workflow.SimpleWorkflow)

	err = s.runPersistence.SetDependenciesState(runId, simpleWf.Dependencies)
	if err != nil {
		return "", err
	}

	go func(wf workflow.Workflow) {
		_, err := wf.Execute(runCtx)
		if err != nil {
			logger.Error("workflow execution failed", "error", err)
		} else {
			logger.Info("workflow execution succeeded")
		}
		if _, err = w.Write(TerminationLogBytes); err != nil {
			ctxLogger.Error("fail to write closing log", "error", err)
		}
		if err = w.Close(); err != nil {
			ctxLogger.Error("fail to close logger", "error", err)
		}
	}(wf)

	go func(simpleWf *workflow.SimpleWorkflow) {
		for state := range simpleWf.StateCh {
			err := s.runPersistence.SetRunState(runId, state)
			if err != nil {
				ctxLogger.Warn("state persistence fail", "runId", runId, "err", err)
			}
		}
	}(simpleWf)

	return runId, nil
}

var _ Selflow = &selflow{}
