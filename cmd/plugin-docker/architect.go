package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-hclog"
	cs "github.com/selflow/selflow/pkg/container-spawner"
	sp "github.com/selflow/selflow/pkg/selflow-plugin"
)

type DockerArchitect struct {
}

func (a *DockerArchitect) ValidateStepConfigSchema(_ context.Context, request *sp.ValidateStepConfigSchema_Request) (*sp.ValidateStepConfigSchema_Response, error) {
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

func (a *DockerArchitect) RunStep(ctx context.Context, request *sp.RunStep_Request) (*sp.RunStep_Response, error) {
	var templateConfig ContainerSpawnConfig

	if err := json.Unmarshal(request.StepConfig, &templateConfig); err != nil {
		return nil, err
	}

	config := cs.SpawnConfig{
		Image:               templateConfig.Image,
		ContainerName:       generateContainerName(),
		ContainerLogsWriter: hclog.Default().StandardWriter(&hclog.StandardLoggerOptions{}),
		Environment:         nil,
		Mounts: []cs.Mountable{
			cs.BinaryMount{
				FileContent: []byte(templateConfig.Commands),
				Destination: "/etc/start",
				ReadOnly:    true,
			},
		},
		Entrypoint: []string{},
	}

	containerExit, err := cs.Spawn(ctx, &config)

	if err != nil {
		return nil, err
	}

	<-containerExit
	return &sp.RunStep_Response{}, nil
}

var _ sp.Architect = &DockerArchitect{}
