package server

import (
	"context"
	"github.com/docker/docker/client"
	"github.com/hashicorp/go-hclog"
	"github.com/selflow/selflow/cmd/selflow-daemon/server/proto"
	"github.com/selflow/selflow/internal/config"
	"github.com/selflow/selflow/pkg/container-spawner/docker"
	"github.com/selflow/selflow/pkg/selflow"
	"github.com/selflow/selflow/pkg/steps/container"
)

type Server struct {
	logger         hclog.Logger
	LogFactory     selflow.LogFactory
	RunPersistence selflow.RunPersistence
	proto.UnimplementedDaemonServer
}

func (s *Server) StartRun(ctx context.Context, request *proto.StartRun_Request) (*proto.StartRun_Response, error) {

	flow, err := config.Parse(request.RunConfig)
	if err != nil {
		return nil, err
	}

	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	workflowBuilder := selflow.WorkflowBuilder{
		StepMappers: []selflow.StepMapper{
			&container.StepMapper{
				ContainerSpawner: docker.NewSpawner(dockerClient),
			},
		},
	}

	self := selflow.NewSelflow(workflowBuilder, s.LogFactory, s.RunPersistence)

	runId, err := self.StartRun(ctx, flow)
	if err != nil {
		return nil, err
	}

	return &proto.StartRun_Response{
		RunId: runId,
	}, nil

}
