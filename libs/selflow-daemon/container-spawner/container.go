package container_spawner

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/ahmetb/dlog"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"io"
	"log"
)

type ContainerSpawnConfig struct {
	Image       string
	Commands    string
	Environment map[string]string
}

type pluginContainer struct {
	containerName string
	containerId   string
}

func pullDockerImage(ctx context.Context, cli client.APIClient, imageName string) error {
	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		return fmt.Errorf("fail to pull image %s : %v", imageName, err)
	}

	_, err = io.Copy(log.Writer(), out)
	if err != nil {
		return fmt.Errorf("fail to pull image %s : %v", imageName, err)
	}

	return nil
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

func createContainer(ctx context.Context, cli client.APIClient, config *SpawnConfig) (*pluginContainer, error) {

	containerConfig := &container.Config{}
	containerConfig.Env = buildEnvMap(config.Environment)
	containerConfig.Image = config.Image
	containerConfig.Entrypoint = config.Entrypoint
	containerConfig.ExposedPorts = nat.PortSet{}

	hostConfig := &container.HostConfig{}
	hostConfig.PortBindings = nat.PortMap{}
	mounts, err := ToMountList(config.Mounts)
	if err != nil {
		return nil, err
	}
	hostConfig.Mounts = mounts

	for _, portForwardConfig := range config.PortForward {
		containerPort, err := nat.NewPort("tcp", portForwardConfig.Container)
		if err != nil {
			return nil, err
		}
		containerConfig.ExposedPorts[containerPort] = struct{}{}
		hostConfig.PortBindings[containerPort] = []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: portForwardConfig.Host,
			},
		}
	}

	networkConfig := &network.NetworkingConfig{}
	networkConfig.EndpointsConfig = map[string]*network.EndpointSettings{}
	for _, networkName := range config.Networks {
		networkConfig.EndpointsConfig[networkName] = &network.EndpointSettings{}
	}

	response, err := cli.ContainerCreate(
		ctx,
		containerConfig,
		hostConfig,
		networkConfig,
		nil,
		config.ContainerName,
	)
	if err != nil {
		if client.IsErrNotFound(err) {
			if err = pullDockerImage(ctx, cli, config.Image); err != nil {
				return nil, err
			}
			response, err = cli.ContainerCreate(
				ctx,
				containerConfig,
				hostConfig,
				networkConfig,
				nil,
				config.ContainerName,
			)

			if err != nil {
				return nil, err
			}

		} else {
			return nil, err
		}
	}

	return &pluginContainer{
		containerName: config.ContainerName,
		containerId:   response.ID,
	}, nil
}

func startContainer(ctx context.Context, cli client.APIClient, config *pluginContainer) error {
	err := cli.ContainerStart(
		ctx,
		config.containerId,
		types.ContainerStartOptions{},
	)

	if err != nil {
		return fmt.Errorf("fail to start container-spawner [%s] : %s", config.containerName, err)
	}
	return nil
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

func transferContainerLogs(ctx context.Context, cli client.APIClient, config *pluginContainer, writer io.Writer) error {
	out, err := cli.ContainerLogs(ctx, config.containerId, types.ContainerLogsOptions{
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
