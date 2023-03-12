package container_spawner

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"io"
	"log"
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

func Spawn(ctx context.Context, config *SpawnConfig) (chan int64, error) {
	containerTerminated := make(chan int64, 1)

	cli, err := GetClient()
	if err != nil {
		return containerTerminated, err
	}

	var ctn *pluginContainer

	if ctn, err = createContainer(ctx, cli, config); err != nil {
		return containerTerminated, err
	}
	if err = startContainer(ctx, cli, ctn); err != nil {
		return containerTerminated, err
	}

	go func() {

		err = transferContainerLogs(ctx, cli, ctn, config.ContainerLogsWriter)
		if err != nil {
			log.Printf("[WARN] Fail to transfer container logs : %v\n", err)
		}

		containerOkBodyCh, _ := cli.ContainerWait(ctx, ctn.containerId, container.WaitConditionNotRunning)
		containerOkBody := <-containerOkBodyCh
		containerTerminated <- containerOkBody.StatusCode
	}()

	return containerTerminated, err

}
