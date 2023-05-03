package server

import (
	"context"
	"github.com/docker/docker/client"
	"github.com/selflow/selflow/cmd/selflow-daemon/server/proto"
	"github.com/selflow/selflow/internal/config"
	"github.com/selflow/selflow/pkg/container-spawner/docker"
	"github.com/selflow/selflow/pkg/steps/container"
	"github.com/selflow/selflow/pkg/workflow"
	"log"
)

type Server struct {
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

	workflowBuilder := WorkflowBuilder{
		stepMappers: []StepMapper{
			&container.StepMapper{
				ContainerSpawner: docker.NewSpawner(dockerClient),
			},
		},
	}

	wf, err := workflowBuilder.BuildWorkflow(flow)
	if err != nil {
		return nil, err
	}

	go func(wf workflow.Workflow) {
		_, err := wf.Execute(ctx)
		if err != nil {
			log.Printf("[ERROR] workflow execution failed : %v", err)
		} else {
			log.Printf("[INFO] workflow execution succeeded")
		}
	}(wf)

	return &proto.StartRun_Response{
		RunId:       "toto",
		Diagnostics: nil,
	}, nil

}
