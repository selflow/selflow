package container_spawner

import (
	"context"
	"io"
)

type PortForwardConfig struct {
	Host      string
	Container string
}

type SpawnConfig struct {
	Image               string
	ContainerName       string
	ContainerLogsWriter io.Writer
	Environment         map[string]string
	Entrypoint          []string
	Mounts              []Mountable
	PortForward         []PortForwardConfig
	Networks            []string
}

func SpawnAsync(ctx context.Context, config *SpawnConfig) (string, error) {
	cli, err := GetClient()
	if err != nil {
		return "", err
	}

	var ctn *pluginContainer

	if ctn, err = createContainer(ctx, cli, config); err != nil {
		return "", err
	}
	if err = startContainer(ctx, cli, ctn); err != nil {
		return "", err
	}

	return ctn.containerId, nil

}
