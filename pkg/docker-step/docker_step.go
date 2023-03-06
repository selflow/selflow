package docker_step

import (
	"context"
	containerSpawner "github.com/selflow/selflow/pkg/container-spawner"
	"github.com/selflow/selflow/pkg/workflow"
	"os"
	"strconv"
)

type DockerStepConfig struct {
	Image    string
	Commands string
}

type DockerStep struct {
	workflow.SimpleStep
	Image    string
	Commands string
}

func (d *DockerStep) Execute(ctx context.Context) (map[string]string, error) {
	config := containerSpawner.SpawnConfig{
		Image:               d.Image,
		ContainerName:       "cn",
		ContainerLogsWriter: os.Stdout,
		Environment:         nil,
		Mounts: []containerSpawner.Mountable{
			containerSpawner.BinaryMount{
				FileContent: []byte(d.Commands),
				Destination: "/etc/start",
				ReadOnly:    true,
			},
		},
		Entrypoint: []string{},
	}

	containerExit, err := containerSpawner.Spawn(ctx, &config)

	if err != nil {
		d.SetStatus(workflow.ERROR)
		return nil, err
	}

	exitCode := <-containerExit
	if exitCode == 0 {
		d.SetStatus(workflow.SUCCESS)
	} else {
		d.SetStatus(workflow.ERROR)
	}
	return map[string]string{"EXIT_CODE": strconv.FormatInt(exitCode, 10)}, nil
}

var _ workflow.Step = &DockerStep{}
