package server

import (
	"context"
	"github.com/docker/docker/client"
	proto2 "github.com/selflow/selflow/apps/selflow-daemon/server/proto"
	"github.com/selflow/selflow/internal/config"
	"github.com/selflow/selflow/pkg/container-spawner/docker"
	"github.com/selflow/selflow/pkg/selflow"
	"github.com/selflow/selflow/pkg/steps/container"
	"log/slog"
)

type Server struct {
	LogFactory     selflow.LogFactory
	RunPersistence selflow.RunPersistence
	proto2.UnimplementedDaemonServer
}

func (s *Server) StartRun(ctx context.Context, request *proto2.StartRun_Request) (*proto2.StartRun_Response, error) {

	ctx = context.Background()
	flow, err := config.Parse(request.RunConfig)
	if err != nil {
		return nil, err
	}

	slog.DebugContext(ctx, "Connecting to Docker daemon")
	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		slog.ErrorContext(ctx, "Fail to connect to Docker Daemon")
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

	return &proto2.StartRun_Response{
		RunId: runId,
	}, nil

}