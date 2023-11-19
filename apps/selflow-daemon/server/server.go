package server

import (
	"context"
	"github.com/docker/docker/client"
	"github.com/selflow/selflow/apps/selflow-daemon/server/proto"
	"github.com/selflow/selflow/libs/core/selflow"
	"github.com/selflow/selflow/libs/selflow-daemon/config"
	"github.com/selflow/selflow/libs/selflow-daemon/container-spawner/docker"
	"github.com/selflow/selflow/libs/selflow-daemon/sfenvironment"
	"github.com/selflow/selflow/libs/selflow-daemon/steps/container"
	"log/slog"
	"path"
)

type Server struct {
	LogFactory     selflow.LogFactory
	RunPersistence selflow.RunPersistence
	proto.UnimplementedDaemonServer
}

func (s *Server) StartRun(ctx context.Context, request *proto.StartRun_Request) (*proto.StartRun_Response, error) {

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
				ContainerSpawner: docker.NewSpawner(dockerClient, path.Join(sfenvironment.GetDaemonBaseDirectory(), "tmp"), path.Join(sfenvironment.GetDaemonHostBaseDirectory(), "tmp")),
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
