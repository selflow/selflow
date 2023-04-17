package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/ahmetb/dlog"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"io"
	"log"
	"os"
	"strings"
)

type dockerSpawner struct {
	dockerClient client.APIClient
}

func buildEnvString(key string, value string) string {
	return fmt.Sprintf("%s=%s", key, value)
}

func buildEnvMap(environmentVariables map[string]string) []string {
	envAsString := make([]string, 0, len(environmentVariables))
	for key, value := range environmentVariables {
		envAsString = append(envAsString, buildEnvString(key, value))
	}
	return envAsString
}

func (d *dockerSpawner) pullDockerImage(ctx context.Context, imageName string) error {
	out, err := d.dockerClient.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		return fmt.Errorf("fail to pull image %s : %v", imageName, err)
	}

	_, err = io.Copy(log.Writer(), out)
	if err != nil {
		return fmt.Errorf("fail to pull image %s : %v", imageName, err)
	}

	return nil
}

func (d *dockerSpawner) createContainer(ctx context.Context, config *ContainerConfig) (string, error) {

	file, err := os.CreateTemp("", "selflow-start")

	if err != nil {
		return "", err
	}

	_, err = file.Write([]byte(config.Commands))
	if err != nil {
		return "", err
	}

	containerConfig := &container.Config{}
	containerConfig.Env = buildEnvMap(config.Environment)
	containerConfig.Image = config.Image
	containerConfig.Entrypoint = strings.Split("/bin/sh /entrypoint.sh", " ")
	containerConfig.ExposedPorts = nat.PortSet{}

	hostConfig := &container.HostConfig{}
	hostConfig.PortBindings = nat.PortMap{}
	hostConfig.Mounts = []mount.Mount{
		{
			Type:     mount.TypeBind,
			Source:   file.Name(),
			ReadOnly: true,
			Target:   "/entrypoint.sh",
		},
	}

	networkConfig := &network.NetworkingConfig{}
	networkConfig.EndpointsConfig = map[string]*network.EndpointSettings{}
	for _, networkName := range config.Networks {
		networkConfig.EndpointsConfig[networkName] = &network.EndpointSettings{}
	}

	response, err := d.dockerClient.ContainerCreate(
		ctx,
		containerConfig,
		hostConfig,
		networkConfig,
		nil,
		config.ContainerName,
	)
	if err != nil {
		if client.IsErrNotFound(err) {
			if err = d.pullDockerImage(ctx, config.Image); err != nil {
				return "", err
			}
			response, err = d.dockerClient.ContainerCreate(
				ctx,
				containerConfig,
				hostConfig,
				networkConfig,
				nil,
				config.ContainerName,
			)

			if err != nil {
				return "", err
			}

		} else {
			return "", err
		}
	}

	return response.ID, nil
}

func (d *dockerSpawner) startContainer(ctx context.Context, containerId string) error {
	err := d.dockerClient.ContainerStart(
		ctx,
		containerId,
		types.ContainerStartOptions{},
	)

	if err != nil {
		return fmt.Errorf("fail to start container-spawner [%s] : %s", containerId, err)
	}
	return nil
}

func (d *dockerSpawner) StartContainerDetached(ctx context.Context, config *ContainerConfig) (string, error) {

	containerId, err := d.createContainer(ctx, config)
	if err != nil {
		return "", err
	}
	err = d.startContainer(ctx, containerId)
	if err != nil {
		return "", err
	}

	return containerId, nil
}

type DockerWriter struct {
	io.Writer
}

func (dw *DockerWriter) Write(p []byte) (n int, err error) {

	parsedReader := dlog.NewReader(bytes.NewReader(p))
	scanner := bufio.NewScanner(parsedReader)

	for scanner.Scan() {
		_, err = dw.Writer.Write(scanner.Bytes())
		if err != nil {
			return 0, err
		}
	}

	return len(p), nil
}

func (d *dockerSpawner) TransferContainerLogs(ctx context.Context, containerId string, writer io.Writer) error {
	out, err := d.dockerClient.ContainerLogs(ctx, containerId, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
	})

	if err != nil {
		return err
	}

	_, err = io.Copy(&DockerWriter{writer}, out)
	if err != nil {
		return err
	}
	return nil
}

func (d *dockerSpawner) WaitContainer(ctx context.Context, containerId string) (int64, error) {
	containerOkBodyCh, errCh := d.dockerClient.ContainerWait(ctx, containerId, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		return -1, err
	case containerOkBody := <-containerOkBodyCh:
		return containerOkBody.StatusCode, nil
	}
}
