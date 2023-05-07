package selflow

import (
	"context"
	"github.com/selflow/selflow/internal/config"
	"github.com/selflow/selflow/pkg/sflog"
	"github.com/selflow/selflow/pkg/workflow"
	"io"
)

type Selflow interface {
	StartRun(ctx context.Context, config *config.Flow) (runId string, err error)
}

type LogFactory interface {
	GetRunLogger(runId string) (io.Reader, io.WriteCloser, error)
}

type selflow struct {
	workflowBuilder WorkflowBuilder
	logFactory      LogFactory
}

func NewSelflow(workflowBuilder WorkflowBuilder, logFactory LogFactory) Selflow {
	return &selflow{
		workflowBuilder: workflowBuilder,
		logFactory:      logFactory,
	}
}

func (s *selflow) StartRun(ctx context.Context, flow *config.Flow) (string, error) {
	runId := "toto"

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

	go func(wf workflow.Workflow) {
		_, err := wf.Execute(runCtx)
		if err != nil {
			logger.Error("workflow execution failed", "error", err)
		} else {
			logger.Info("workflow execution succeeded")
		}
		err = w.Close()
		if err != nil {
			ctxLogger.Error("fail to close logger", "error", err)
		}
	}(wf)

	return runId, nil
}

var _ Selflow = &selflow{}
