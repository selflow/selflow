package runners

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/selflow/selflow/internal/config"
	cs "github.com/selflow/selflow/pkg/container-spawner"
	sp "github.com/selflow/selflow/pkg/selflow-plugin"
	"os"
)

func startRunnerContainer(ctx context.Context, flow config.Flow, runId string, logger hclog.Logger) (chan int64, error) {
	//flowFileName, err := getFlowFileName(flow)
	flowAsBytes, err := json.Marshal(flow)
	if err != nil {
		return nil, err
	}

	spawnConfig := &cs.SpawnConfig{}
	spawnConfig.Image = "selflow-runner"
	spawnConfig.ContainerLogsWriter = logger.StandardWriter(&hclog.StandardLoggerOptions{
		InferLevels: true,
	})
	spawnConfig.Environment = map[string]string{
		sp.Handshake.MagicCookieKey: sp.Handshake.MagicCookieValue,
	}
	spawnConfig.ContainerName = fmt.Sprintf("%s-%s", runnerContainerBaseName, runId)
	spawnConfig.Mounts = []cs.Mountable{
		cs.BinaryMount{
			FileContent:   flowAsBytes,
			Destination:   "/etc/selflow/config.json",
			ReadOnly:      true,
			TempDirectory: os.Getenv("TMP_FILE_HOST_DIR"),
		},
	}
	spawnConfig.Networks = []string{
		daemonNetwork,
	}

	if debugPort := ctx.Value("debug-port"); debugPort != nil {
		spawnConfig.PortForward = []cs.PortForwardConfig{
			{
				Host:      debugPort.(string),
				Container: debugPort.(string), // Default port with delve
			},
		}
	}

	exitCodeCh, err := cs.Spawn(ctx, spawnConfig)
	if err != nil {
		return nil, err
	}

	return exitCodeCh, nil
}
