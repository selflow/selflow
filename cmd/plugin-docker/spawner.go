package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/hashicorp/go-hclog"
	"io"
	"log"
	"os"
	"strconv"
	"time"
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

func pullDockerImage(ctx context.Context, cli *client.Client, imageName string) error {
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

func generateContainerName() string {
	return strconv.FormatInt(time.Now().UnixMicro(), 36)
}

func createContainer(ctx context.Context, cli *client.Client, config *ContainerSpawnConfig) (*pluginContainer, error) {
	containerName := generateContainerName()

	tmpFile, err := os.CreateTemp("", "shell")
	if err != nil {
		return nil, err
	}

	_, err = tmpFile.WriteString(config.Commands)
	if err != nil {
		return nil, err
	}

	response, err := cli.ContainerCreate(
		ctx,
		&container.Config{
			Env:   buildEnvMap(config.Environment),
			Image: config.Image,
			Entrypoint: []string{
				"/bin/sh",
				"/etc/start",
			},
		},
		&container.HostConfig{
			//AutoRemove: true,
			Mounts: []mount.Mount{
				{
					Type:     mount.TypeBind,
					Source:   tmpFile.Name(),
					Target:   "/etc/start",
					ReadOnly: true,
				},
			},
		},
		&network.NetworkingConfig{EndpointsConfig: map[string]*network.EndpointSettings{}},
		nil,
		containerName,
	)

	if err != nil {
		return nil, fmt.Errorf("fail to create container : %v", err)
	}
	return &pluginContainer{
		containerName: containerName,
		containerId:   response.ID,
	}, nil
}

func startContainer(ctx context.Context, cli *client.Client, config *pluginContainer) error {
	err := cli.ContainerStart(
		ctx,
		config.containerId,
		types.ContainerStartOptions{},
	)

	if err != nil {
		return fmt.Errorf("fail to start container [%s] : %s", config.containerName, err)
	}
	return nil
}

func transferContainerLogs(ctx context.Context, cli *client.Client, config *pluginContainer, writer io.Writer) error {
	out, err := cli.ContainerLogs(ctx, config.containerId, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Timestamps: false,
		Follow:     true,
		Tail:       "",
		Details:    false,
	})

	if err != nil {
		return fmt.Errorf("fail to read plugin logs : %v", err)
	}

	_, err = io.Copy(writer, out)
	if err != nil {
		return fmt.Errorf("fail to transfer plugin logs: %v", err)
	}
	return nil
}

func Spawn(ctx context.Context, config *ContainerSpawnConfig) error {

	cli, err := GetClient()
	if err != nil {
		return err
	}

	var ctn *pluginContainer

	// TODO : check image exists
	if err = pullDockerImage(ctx, cli, config.Image); err != nil {
		return err
	}
	if ctn, err = createContainer(ctx, cli, config); err != nil {
		return err
	}
	if err = startContainer(ctx, cli, ctn); err != nil {
		return err
	}

	pluginLogger := hclog.Default().Named(fmt.Sprintf("plugin-%s", ctn.containerName))

	err = transferContainerLogs(ctx, cli, ctn, pluginLogger.StandardWriter(&hclog.StandardLoggerOptions{}))
	if err != nil {
		pluginLogger.Warn("fail to register plugin logs", "err", err)
	}

	cli.ContainerWait(ctx, ctn.containerId, container.WaitConditionNotRunning)
	return nil

}
