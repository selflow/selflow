package main

import (
	"context"
	"github.com/docker/docker/client"
	cs "github.com/selflow/selflow/libs/selflow-daemon/container-spawner"
)

type ContainerSpawner interface {
	client.APIClient
	SpawnAsync(ctx context.Context, config *cs.SpawnConfig) (string, error)
}

type containerSpawner struct {
	client.APIClient
}

func (i containerSpawner) SpawnAsync(ctx context.Context, config *cs.SpawnConfig) (string, error) {
	return cs.SpawnAsync(ctx, config)
}
