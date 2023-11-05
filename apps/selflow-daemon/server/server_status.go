package server

import (
	"context"
	"github.com/selflow/selflow/apps/selflow-daemon/server/proto"
	"github.com/selflow/selflow/pkg/workflow"
)

func runStateToProtoState(state map[string]workflow.Status) map[string]*proto.GetRunStatus_Status {
	responseState := map[string]*proto.GetRunStatus_Status{}

	for stepId, status := range state {
		responseState[stepId] = &proto.GetRunStatus_Status{
			Name:          status.GetName(),
			Code:          int32(status.GetCode()),
			IsFinished:    status.IsFinished(),
			IsCancellable: status.IsCancellable(),
		}
	}
	return responseState
}

func runDependenciesToProtoDependencies(dependencies map[string][]string) map[string]*proto.GetRunStatus_Dependence {
	protoDependencies := map[string]*proto.GetRunStatus_Dependence{}

	for stepId, stepDependencies := range dependencies {
		protoDependencies[stepId] = &proto.GetRunStatus_Dependence{Dependencies: stepDependencies}
	}

	return protoDependencies
}

func (s *Server) GetRunStatus(_ context.Context, req *proto.GetRunStatus_Request) (*proto.GetRunStatus_Response, error) {
	state, err := s.RunPersistence.GetRunState(req.GetRunId())
	if err != nil {
		return nil, err
	}

	dependencies, err := s.RunPersistence.GetRunDependencies(req.GetRunId())
	if err != nil {
		return nil, err
	}

	stepDefinitions, err := s.RunPersistence.GetRunStepDefinitions(req.GetRunId())
	if err != nil {
		return nil, err
	}

	return &proto.GetRunStatus_Response{
		State:           runStateToProtoState(state),
		Dependencies:    runDependenciesToProtoDependencies(dependencies),
		StepDefinitions: stepDefinitions,
	}, nil
}
