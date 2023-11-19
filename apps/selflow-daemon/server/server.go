package server

import (
	"context"
	"github.com/docker/docker/client"
	"github.com/google/uuid"
	"github.com/selflow/selflow/apps/selflow-daemon/server/proto"
	"github.com/selflow/selflow/libs/core/selflow"
	"github.com/selflow/selflow/libs/core/sflog"
	"github.com/selflow/selflow/libs/core/workflow"
	"github.com/selflow/selflow/libs/selflow-daemon/config"
	"github.com/selflow/selflow/libs/selflow-daemon/container-spawner/docker"
	"github.com/selflow/selflow/libs/selflow-daemon/sfenvironment"
	"github.com/selflow/selflow/libs/selflow-daemon/steps/container"
	"io"
	"log/slog"
	"path"
)

type LogFactory interface {
	GetRunLogger(runId string) (io.Reader, io.WriteCloser, error)
	GetRunReader(runId string) (chan string, error)
}

type Server struct {
	LogFactory     LogFactory
	RunPersistence RunPersistence
	proto.UnimplementedDaemonServer
}

type RunPersistence interface {
	selflow.WorkflowEventHandler

	GetRunState(runId string) (map[string]workflow.Status, error)
	GetRunDependencies(runId string) (map[string][]string, error)
	GetRunStepDefinitions(runId string) (map[string][]byte, error)
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

	self := selflow.NewSelflow(workflowBuilder, s.RunPersistence)

	runId := uuid.NewString()

	_, w, err := s.LogFactory.GetRunLogger(runId)
	if err != nil {
		slog.ErrorContext(ctx, "Fail to get run logger", "error", err)
		return nil, err
	}

	logger := sflog.NewLoggerWithWriter(ctx, w)
	ctx = sflog.ResetContextLogHandler(ctx, logger.Handler())

	go func() {
		ctx = sflog.ResetContextLogHandler(ctx, slog.NewJSONHandler(w, &slog.HandlerOptions{Level: slog.LevelDebug}))

		defer func() {
			// Close logger at the end of the run
			if _, err = w.Write(selflow.TerminationLogBytes); err != nil {
				slog.ErrorContext(ctx, "fail to write closing log", "error", err)
			}
		}()

		// Start selflow run
		err := self.StartRunWithId(ctx, flow, runId)
		if err != nil {
			slog.ErrorContext(ctx, "Selflow exited with an error", "error", err)
			return
		}

	}()

	return &proto.StartRun_Response{
		RunId: runId,
	}, nil

}
