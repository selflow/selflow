package main

import (
	"context"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/selflow/selflow/internal/sfenvironment"
	cs "github.com/selflow/selflow/pkg/container-spawner"
	"net/url"
)

func (sc *selflowClient) createNetworkIfNotExists(ctx context.Context, networkName string) (string, error) {
	network, err := sc.dockerClient.NetworkInspect(ctx, networkName, types.NetworkInspectOptions{})
	if err != nil {
		if !client.IsErrNotFound(err) {
			return "", nil
		}

		networkCreateResponse, err := sc.dockerClient.NetworkCreate(ctx, networkName, types.NetworkCreate{})
		if err != nil {
			return "", err
		}

		return networkCreateResponse.ID, nil

	}

	return network.ID, nil
}

func (sc *selflowClient) clearContainer(ctx context.Context, containerName string) error {

	err := sc.dockerClient.ContainerStop(ctx, containerName, nil)
	if err != nil {
		if client.IsErrNotFound(err) {
			return nil
		}
		return err
	}

	return sc.dockerClient.ContainerRemove(ctx, containerName, types.ContainerRemoveOptions{
		RemoveVolumes: true,
		Force:         false,
	})
}

func (sc *selflowClient) startDaemon(ctx context.Context, selflowClient *selflowClient) (string, error) {
	var containerId string

	containerId, err := sc.getRunningDaemon(ctx, selflowClient.daemonName)
	if err != nil {
		if !errors.Is(err, ContainerNotFound) {
			return "", err
		}

		err = sc.clearContainer(ctx, selflowClient.daemonName)
		if err != nil && !client.IsErrNotFound(err) {
			return "", err
		}
		containerId, err = sc.createDaemon(ctx, selflowClient.daemonName)
		if err != nil {
			return "", err
		}
	}

	return containerId, nil
}

func (sc *selflowClient) getRunningDaemon(ctx context.Context, containerName string) (string, error) {

	container, err := sc.dockerClient.ContainerInspect(ctx, containerName)
	if err != nil {
		return "", ContainerNotFound
	}

	return container.ID, nil
}

func (sc *selflowClient) createDaemon(ctx context.Context, containerName string) (string, error) {

	dockerHostUrl, _ := url.Parse(sc.dockerClient.DaemonHost())

	portForwardConfig := []cs.PortForwardConfig{
		{
			Host:      sc.daemonPort,
			Container: "10011",
		},
	}

	if sc.daemonIsDebug {
		portForwardConfig = append(portForwardConfig, cs.PortForwardConfig{
			Host:      sc.daemonDebugPort,
			Container: sc.daemonDebugPort,
		})
	}
	return cs.SpawnAsync(ctx, &cs.SpawnConfig{
		Image:         "selflow-daemon",
		ContainerName: containerName,
		Entrypoint:    nil,
		Environment: map[string]string{
			sfenvironment.DaemonPortEnvKey:              sc.daemonPort,
			sfenvironment.DaemonDebugPortEnvKey:         sc.daemonDebugPort,
			sfenvironment.DaemonNameEnvKey:              sc.daemonName,
			sfenvironment.DaemonBaseDirectoryEnvKey:     sc.daemonBaseDirectory,
			sfenvironment.DaemonNetworkEnvKey:           sc.daemonNetworkName,
			sfenvironment.DaemonImageEnvKey:             sc.daemonDockerImage,
			sfenvironment.DaemonHostBaseDirectoryEnvKey: sc.daemonHostBaseDirectory,
		},
		Mounts: []cs.Mountable{
			cs.FileMount{
				SourceFileName: dockerHostUrl.Path,
				Destination:    "/var/run/docker.sock",
				ReadOnly:       false,
			},
			cs.FileMount{
				SourceFileName: sc.daemonHostBaseDirectory,
				Destination:    sc.daemonBaseDirectory,
			},
		},
		PortForward: portForwardConfig,
		Networks:    []string{containerName},
	})
}
