package main

import (
	"context"
	"io"
)

type ContainerConfig struct {
	Image               string
	ContainerName       string
	Commands            string
	ContainerLogsWriter io.Writer
	Environment         map[string]string
	Entrypoint          string
	Mounts              []Mountable
	OpenPorts           []OpenPortConfig
	Networks            []string
}

type OpenPortConfig struct {
	ContainerPort string
	OutsidePort   string
}

type Mount struct {
	ArtifactName string
	// Destination is the absolute path where the artifact should be stored
	Destination string
}

type Mountable interface {
	ToMount() (Mount, error)
}

type BinaryMount struct {
	FileContent        []byte
	Destination        string
	ReadOnly           bool
	HostTempDirectory  string
	LocalTempDirectory string
}

type ContainerSpawner interface {
	StartContainerDetached(ctx context.Context, config *ContainerConfig) (string, error)
	TransferContainerLogs(ctx context.Context, containerId string, writer io.Writer) error
	WaitContainer(ctx context.Context, containerId string) (int64, error)
}
