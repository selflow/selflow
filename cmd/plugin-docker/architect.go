package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	sp "github.com/selflow/selflow/pkg/selflow-plugin"
)

type DockerArchitect struct {
}

func (a *DockerArchitect) ValidateStepConfigSchema(request *sp.ValidateStepConfigSchema_Request) (*sp.ValidateStepConfigSchema_Response, error) {
	var v = validator.New()

	configAsBytes := request.GetStepConfig()
	var conf ContainerSpawnConfig

	if err := json.Unmarshal(configAsBytes, &conf); err != nil {
		return &sp.ValidateStepConfigSchema_Response{
			Valid: false,
			Diagnotics: []*sp.Diagnostic{
				{
					Type:    sp.DiagnosticType_ERROR,
					Message: fmt.Sprintf("%v", err),
				},
			},
		}, nil
	}

	if err := v.Struct(conf); err != nil {
		return &sp.ValidateStepConfigSchema_Response{
			Valid: false,
			Diagnotics: []*sp.Diagnostic{
				{
					Type:    sp.DiagnosticType_ERROR,
					Message: fmt.Sprintf("%v", err),
				},
			},
		}, nil
	}
	return &sp.ValidateStepConfigSchema_Response{
		Valid: true,
	}, nil

}

func (a *DockerArchitect) RunStep(request *sp.RunStep_Request) (*sp.RunStep_Response, error) {
	var config ContainerSpawnConfig

	if err := json.Unmarshal(request.StepConfig, &config); err != nil {
		return nil, err
	}

	if err := Spawn(context.Background(), &config); err != nil {
		return nil, err
	}

	return &sp.RunStep_Response{}, nil
}

var _ sp.Architect = &DockerArchitect{}
