package runners

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-hclog"
	cs "github.com/selflow/selflow/pkg/container-spawner"
	"os"
)

type containerSpawner struct {
	runId string
}

func (c containerSpawner) SpawnContainer(ctx context.Context, containerId string, environmentVariables map[string]string, cmd string, image string) error {

	tmpFile, err := os.CreateTemp("/etc/selflow", "start-")
	if err != nil {
		return err
	}

	_, err = tmpFile.WriteString(cmd)
	if err != nil {
		return err
	}

	conf := &cs.SpawnConfig{}

	containerNameSuffix := fmt.Sprintf("cc-%s-%s", c.runId, containerId)

	conf.ContainerLogsWriter = hclog.Default().Named(containerNameSuffix).StandardWriter(&hclog.StandardLoggerOptions{
		ForceLevel: hclog.Debug,
	})
	conf.Image = image
	conf.Mounts = []cs.Mountable{
		cs.BinaryMount{
			FileContent:   []byte(cmd),
			Destination:   "/start.sh",
			ReadOnly:      true,
			TempDirectory: os.Getenv("TMP_FILE_HOST_DIR"),
		},
	}
	conf.Entrypoint = []string{"/bin/sh", "/start.sh"}
	conf.Environment = environmentVariables

	ch, err := cs.Spawn(ctx, conf)
	if err != nil {
		return err
	}
	<-ch
	return nil
}
