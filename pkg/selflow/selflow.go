package selflow

import (
	"context"
	"github.com/selflow/selflow/internal/config"
	"github.com/selflow/selflow/pkg/workflow"
	"log"
)

type Selflow interface {
	StartRun(ctx context.Context, config *config.Flow) (runId string, err error)
}

type selflow struct {
	workflowBuilder WorkflowBuilder
}

func NewSelflow(workflowBuilder WorkflowBuilder) Selflow {
	return &selflow{
		workflowBuilder: workflowBuilder,
	}
}

func (s *selflow) StartRun(ctx context.Context, flow *config.Flow) (runId string, err error) {
	wf, err := s.workflowBuilder.BuildWorkflow(flow)
	if err != nil {
		return "", err
	}

	go func(wf workflow.Workflow) {
		_, err := wf.Execute(ctx)
		if err != nil {
			log.Printf("[ERROR] workflow execution failed : %v", err)
		} else {
			log.Printf("[INFO] workflow execution succeeded")
		}
	}(wf)

	return "toto", nil
}

var _ Selflow = &selflow{}
