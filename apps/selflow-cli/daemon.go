package main

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/selflow/selflow/libs/selflow-daemon/container-spawner"
	"github.com/selflow/selflow/libs/selflow-daemon/sfenvironment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/url"
	"time"
)

func (sc *selflowClient) createNetworkIfNotExists(ctx context.Context, networkName string) (string, error) {
	network, err := sc.dockerClient.NetworkInspect(ctx, networkName, types.NetworkInspectOptions{})
	if err != nil {
		if !client.IsErrNotFound(err) {
			return "", err
		}

		networkCreateResponse, err := sc.dockerClient.NetworkCreate(ctx, networkName, types.NetworkCreate{})
		if err != nil {
			return "", err
		}

		return networkCreateResponse.ID, nil

	}

	return network.ID, nil
}

func (sc *selflowClient) clearDaemon(ctx context.Context) error {

	err := sc.dockerClient.ContainerStop(ctx, sc.daemonName, container.StopOptions{})
	if err != nil {
		if client.IsErrNotFound(err) {
			return nil
		}
		return err
	}

	return sc.dockerClient.ContainerRemove(ctx, sc.daemonName, types.ContainerRemoveOptions{
		RemoveVolumes: true,
		Force:         false,
	})
}

func (sc *selflowClient) waitDaemon(ctx context.Context) error {
	var err error
	for i := 0; i < 5; i++ {
		_, err = grpc.Dial("localhost:10011", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err == nil {
			return nil
		}
		time.Sleep(2 * time.Second)
	}

	return sc.waitDaemon(ctx)
}

func (sc *selflowClient) startDaemon(ctx context.Context) (string, error) {
	var containerId string

	containerId, err := sc.getRunningDaemon(ctx)
	if err != nil {
		err = sc.clearDaemon(ctx)
		if err != nil && !client.IsErrNotFound(err) {
			return "", err
		}
		containerId, err = sc.createDaemon(ctx)
		if err != nil {
			return "", err
		}
	}

	return containerId, nil
}

func (sc *selflowClient) getRunningDaemon(ctx context.Context) (string, error) {

	container, err := sc.dockerClient.ContainerInspect(ctx, sc.daemonName)
	if err != nil {
		return "", ContainerNotFound
	}

	return container.ID, nil
}

func (sc *selflowClient) createDaemon(ctx context.Context) (string, error) {

	dockerHostUrl, _ := url.Parse(sc.dockerClient.DaemonHost())

	portForwardConfig := []container_spawner.PortForwardConfig{
		{
			Host:      sc.daemonPort,
			Container: "10011",
		},
	}

	if sc.daemonIsDebug {
		portForwardConfig = append(portForwardConfig, container_spawner.PortForwardConfig{
			Host:      sc.daemonDebugPort,
			Container: sc.daemonDebugPort,
		})
	}
	return sc.dockerClient.SpawnAsync(ctx, &container_spawner.SpawnConfig{
		Image:         "selflow-daemon",
		ContainerName: sc.daemonName,
		Entrypoint:    nil,
		Environment: map[string]string{
			sfenvironment.DaemonPortEnvKey:              sc.daemonPort,
			sfenvironment.DaemonDebugPortEnvKey:         sc.daemonDebugPort,
			sfenvironment.DaemonNameEnvKey:              sc.daemonName,
			sfenvironment.DaemonBaseDirectoryEnvKey:     sc.daemonBaseDirectory,
			sfenvironment.DaemonNetworkEnvKey:           sc.daemonNetworkName,
			sfenvironment.DaemonHostBaseDirectoryEnvKey: sc.daemonHostBaseDirectory,
			sfenvironment.UseJsonLogEnvKey:              "TRUE",
			sfenvironment.LogLevelEnvKey:                "DEBUG",
		},
		Mounts: []container_spawner.Mountable{
			container_spawner.FileMount{
				SourceFileName: dockerHostUrl.Path,
				Destination:    "/var/run/docker.sock",
				ReadOnly:       false,
			},
			container_spawner.FileMount{
				SourceFileName: sc.daemonHostBaseDirectory,
				Destination:    sc.daemonBaseDirectory,
			},
		},
		PortForward: portForwardConfig,
		Networks:    []string{sc.daemonNetworkName},
	})
}
