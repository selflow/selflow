package main

import (
	"encoding/json"
	"github.com/selflow/selflow/internal/config"
	dockerStep "github.com/selflow/selflow/pkg/docker-step"
	"github.com/selflow/selflow/pkg/workflow"
	"reflect"
	"testing"
)

func Test_buildWorkflow_classicWorkflow(t *testing.T) {

	conf := &config.Flow{
		Workflow: config.Workflow{
			Steps: map[string]config.StepDefinition{
				"step-a": {
					Kind:  "docker",
					With:  map[string]interface{}{},
					Needs: []string{"step-b"},
				},
				"step-b": {
					Kind: "docker",
					With: map[string]interface{}{},
				},
			},
		},
	}

	expected := workflow.NewWorkflow(1)
	stepA, err := dockerStep.NewDockerStep("step-a", config.StepDefinition{
		Kind: "docker",
		With: map[string]interface{}{},
	})
	stepB, err := dockerStep.NewDockerStep("step-b", config.StepDefinition{
		Kind: "docker",
		With: map[string]interface{}{},
	})
	if err != nil {
		t.Errorf("buildWorkflow() unexpected error : %v", err)
	}

	err = expected.AddStep(stepA, []workflow.Step{stepB})
	if err != nil {
		t.Errorf("buildWorkflow() unexpected error : %v", err)
	}

	err = expected.AddStep(stepB, []workflow.Step{})
	if err != nil {
		t.Errorf("buildWorkflow() unexpected error : %v", err)
	}

	got, err := buildWorkflow(conf)
	if err != nil {
		t.Errorf("buildWorkflow() error = %v, wantErr false", err)
		return
	}

	gotAsJson, _ := json.Marshal(got)
	expectedAsJson, _ := json.Marshal(expected)

	if !reflect.DeepEqual(gotAsJson, expectedAsJson) {
		t.Errorf("buildWorkflow() got = %v, want %v", gotAsJson, expectedAsJson)
	}
}
