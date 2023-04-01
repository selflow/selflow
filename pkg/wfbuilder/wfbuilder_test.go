package wfbuilder

import (
	"errors"
	"github.com/selflow/selflow/internal/config"
	"github.com/selflow/selflow/pkg/workflow"
	"testing"
)

func mockStepMapper(id string, _ config.StepDefinition) (workflow.Step, error) {
	return &workflow.SimpleStep{Id: id, Status: workflow.CREATED}, nil
}

func mockFailingStepMapper(_ string, _ config.StepDefinition) (workflow.Step, error) {
	return nil, errors.New("failing step mapper")
}

const (
	successKind = "success"
	failingKind = "failing"
)

var (
	stepDefA = config.StepDefinition{Kind: successKind}
	stepA    = &workflow.SimpleStep{Id: "a", Status: workflow.CREATED}
	stepDefB = config.StepDefinition{Kind: successKind}
	stepB    = &workflow.SimpleStep{Id: "b", Status: workflow.CREATED}
	stepDefC = config.StepDefinition{Kind: successKind, Needs: []string{"a", "b"}}
	stepC    = &workflow.SimpleStep{Id: "c", Status: workflow.CREATED}
)

func TestBuildWorkflow(t *testing.T) {
	builder := Builder{StepBuilderMap: map[string]StepMapper{
		successKind: mockStepMapper,
		failingKind: mockFailingStepMapper,
	}}

	parsedConfig := &config.Flow{
		Workflow: config.Workflow{
			Steps: map[string]config.StepDefinition{
				"a": stepDefA,
				"b": stepDefB,
				"c": stepDefC,
			},
		},
	}

	expectedWorkflow := workflow.NewWorkflow(3)
	_ = expectedWorkflow.AddStep(stepA, []workflow.Step{})
	_ = expectedWorkflow.AddStep(stepB, []workflow.Step{})
	_ = expectedWorkflow.AddStep(stepC, []workflow.Step{stepA, stepB})

	wf, err := builder.BuildWorkflow(parsedConfig)
	if err != nil {
		t.Errorf("BuildWorkflow() error = %v", err)
	}

	if !wf.Equals(expectedWorkflow) {
		t.Errorf("BuildWorkflow() got = %v, want %v", wf, expectedWorkflow)
	}
}

func TestBuildWorkflow_InvalidStep(t *testing.T) {
	builder := Builder{StepBuilderMap: map[string]StepMapper{
		failingKind: mockFailingStepMapper,
	}}

	parsedConfig := &config.Flow{
		Workflow: config.Workflow{
			Steps: map[string]config.StepDefinition{
				"a": {Kind: "I do not exists"},
			},
		},
	}

	_, err := builder.BuildWorkflow(parsedConfig)
	if err == nil {
		t.Errorf("BuildWorkflow() should throw an error")
	}
}

func TestBuildWorkflow_MissingDependency(t *testing.T) {
	builder := Builder{StepBuilderMap: map[string]StepMapper{
		successKind: mockStepMapper,
	}}

	parsedConfig := &config.Flow{
		Workflow: config.Workflow{
			Steps: map[string]config.StepDefinition{
				"a": {Kind: successKind, Needs: []string{"w"}},
			},
		},
	}

	_, err := builder.BuildWorkflow(parsedConfig)
	if err == nil {
		t.Errorf("BuildWorkflow() should throw an error")
	}
}
