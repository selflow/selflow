package selflow

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/selflow/selflow/libs/core/sflog"
	"github.com/selflow/selflow/libs/core/workflow"
	"github.com/selflow/selflow/libs/selflow-daemon/config"
	"log/slog"
)

const TerminationLogText = "===EOF==="

var TerminationLogBytes = []byte(fmt.Sprintf("%s\n", TerminationLogText))

type Selflow interface {
	StartRun(ctx context.Context, flow *config.Flow) (runId string, err error)
	StartRunWithId(ctx context.Context, flow *config.Flow, id string) (err error)
}

type WorkflowEventHandler interface {
	SetRunState(runId string, state map[string]workflow.Status) error
	SetDependenciesState(runId string, dependencies map[workflow.Step][]workflow.Step) error
	SetRunStepDefinitions(runId string, stepDefinitions map[string]config.StepDefinition) error
}

type selflow struct {
	workflowBuilder WorkflowBuilder
	runPersistence  WorkflowEventHandler
}

func NewSelflow(workflowBuilder WorkflowBuilder, runPersistence WorkflowEventHandler) Selflow {
	return &selflow{
		workflowBuilder: workflowBuilder,
		runPersistence:  runPersistence,
	}
}
func (s *selflow) StartRunWithId(ctx context.Context, flow *config.Flow, runId string) error {

	ctx = sflog.AddArgsToContextLogger(ctx, slog.String("workflowRun", runId))
	slog.DebugContext(ctx, "Building workflow")

	wf, err := s.workflowBuilder.BuildWorkflow(flow)
	if err != nil {
		slog.ErrorContext(ctx, "Fail to build workflow", "error", err)
		return err
	}

	simpleWf := wf.(*workflow.SimpleWorkflow)

	err = s.runPersistence.SetRunStepDefinitions(runId, flow.Workflow.Steps)
	slog.ErrorContext(ctx, "Fail to set workflow steps", "error", err)
	if err != nil {
		return err
	}

	err = s.runPersistence.SetDependenciesState(runId, simpleWf.Dependencies)
	if err != nil {
		slog.ErrorContext(ctx, "Fail to set workflow dependency graph", "error", err)
		return err
	}

	// Watch state changes
	go func(simpleWf *workflow.SimpleWorkflow) {
		for state := range simpleWf.StateCh {
			err := s.runPersistence.SetRunState(runId, state)
			if err != nil {
				slog.WarnContext(ctx, "fail to save run state", "error", err)
			}
		}
	}(simpleWf)

	// Start workflow
	slog.InfoContext(ctx, "Start workflow execution")
	_, err = wf.Execute(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Workflow execution failed", "error", err)
	} else {
		slog.InfoContext(ctx, "Workflow execution succeeded")
	}

	return nil
}

func (s *selflow) StartRun(ctx context.Context, flow *config.Flow) (string, error) {
	runId := uuid.NewString()
	return runId, s.StartRunWithId(ctx, flow, runId)
}

var _ Selflow = &selflow{}
