package selflow_runner_proto

import "context"

type ContainerSpawner interface {
	SpawnContainer(ctx context.Context, containerId string, environmentVariables map[string]string, cmd string, image string) error
}

type GRPCContainerSpawnerClient struct {
	client ContainerSpawnerClient
}

func (c *GRPCContainerSpawnerClient) SpawnContainer(ctx context.Context, containerId string, environmentVariables map[string]string, cmd string, image string) error {
	_, err := c.client.SpawnContainer(ctx, &SpawnContainer_Request{ContainerId: containerId, Environment: environmentVariables, Cmd: cmd, Image: image})
	return err
}

type GRPCContainerSpawnerServer struct {
	Impl ContainerSpawner
	UnimplementedContainerSpawnerServer
}

func (c *GRPCContainerSpawnerServer) SpawnContainer(ctx context.Context, req *SpawnContainer_Request) (*SpawnContainer_Response, error) {
	err := c.Impl.SpawnContainer(ctx, req.ContainerId, req.Environment, req.Cmd, req.Image)
	if err != nil {
		return nil, err
	}
	return &SpawnContainer_Response{}, nil
}
