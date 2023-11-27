package docker

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/ahmetb/dlog"
	"github.com/docker/docker/api/types"
	dockerContainer "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/selflow/selflow/libs/core/workflow"
	"github.com/selflow/selflow/libs/selflow-daemon/steps/container"
	"io"
	"log"
	"log/slog"
	"os"
	"path"
	"strings"
)

const SelflowLabel = "selflow"
const SelflowRunIdLabel = "selflow.runId"

type spawner struct {
	dockerClient client.APIClient
	// tmpDirectory is the directory for entrypoints binding in the process file system
	tmpDirectory string
	// hostTmpDirectory is the directory for entrypoints binding in the host file system..
	// unless the program is running in a docker container, this should be the same as tmpDirectory
	hostTmpDirectory string
}

func NewSpawner(dockerClient client.APIClient, tmpDirectory string, hostTmpDirectory string) container.ContainerSpawner {
	return &spawner{
		dockerClient:     dockerClient,
		tmpDirectory:     tmpDirectory,
		hostTmpDirectory: hostTmpDirectory,
	}
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

func (d *spawner) pullDockerImage(ctx context.Context, imageName string) error {
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

func (d *spawner) createEntrypointFileBinding(config *container.ContainerConfig) (string, error) {
	err := os.MkdirAll(d.tmpDirectory, 0777)
	if err != nil {
		return "", err
	}
	file, err := os.CreateTemp(d.tmpDirectory, "selflow-start")

	if err != nil {
		return "", err
	}

	_, err = file.Write([]byte(config.Commands))
	if err != nil {
		return "", err
	}
	return path.Join(d.hostTmpDirectory, path.Base(file.Name())), nil
}

func (d *spawner) createContainer(ctx context.Context, config *container.ContainerConfig) (string, error) {

	fileName, err := d.createEntrypointFileBinding(config)
	if err != nil {
		slog.ErrorContext(ctx, "Fail to create entrypoint file", "error", err)
		return "", err
	}

	containerConfig := &dockerContainer.Config{}
	containerConfig.Env = buildEnvMap(config.Environment)
	containerConfig.Image = config.Image
	containerConfig.Entrypoint = strings.Split("/bin/sh /entrypoint.sh", " ")
	containerConfig.ExposedPorts = nat.PortSet{}

	hostConfig := &dockerContainer.HostConfig{}
	hostConfig.PortBindings = nat.PortMap{}
	hostConfig.Mounts = []mount.Mount{
		{
			Type:     mount.TypeBind,
			Source:   fileName,
			ReadOnly: true,
			Target:   "/entrypoint.sh",
		},
	}

	defaultMountLabels := map[string]string{
		SelflowLabel: "true",
	}

	if runId, ok := ctx.Value(workflow.RunIdContextKey{}).(string); ok {
		defaultMountLabels[SelflowRunIdLabel] = runId
	}

	for _, mountable := range config.Mounts {
		m, err := mountable.ToMount()
		if err != nil {
			return "", err
		}

		hostConfig.Mounts = append(hostConfig.Mounts, mount.Mount{
			Type:     mount.TypeVolume,
			Source:   m.ArtifactName,
			Target:   m.Destination,
			ReadOnly: false,
			VolumeOptions: &mount.VolumeOptions{
				Labels: defaultMountLabels,
			},
		})
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

func (d *spawner) startContainer(ctx context.Context, containerId string) error {
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

func (d *spawner) StartContainerDetached(ctx context.Context, config *container.ContainerConfig) (string, error) {
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

func (d *spawner) TransferContainerLogs(ctx context.Context, containerId string, writer io.Writer) error {
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

func (d *spawner) WaitContainer(ctx context.Context, containerId string) (int64, error) {
	containerOkBodyCh, errCh := d.dockerClient.ContainerWait(ctx, containerId, dockerContainer.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		return -1, err
	case containerOkBody := <-containerOkBodyCh:
		return containerOkBody.StatusCode, nil
	}
}
