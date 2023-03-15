package runners

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	selflowRunnerProto "github.com/selflow/selflow/internal/selflow-runner-proto"
	sp "github.com/selflow/selflow/pkg/selflow-plugin"
	"google.golang.org/grpc"
	"log"
	"net"
)

func initContainer(ctx context.Context, runId string) error {
	host := fmt.Sprintf("%s-%s", runnerContainerBaseName, runId)
	tcpAddr := fmt.Sprintf("%s:11001", host)
	log.Printf("[DEBUG] initializing connection with runner %s on %s", runId, tcpAddr)

	addr, err := net.ResolveTCPAddr("tcp", tcpAddr)
	if err != nil {
		return err
	}

	pluginClient := plugin.NewClient(&plugin.ClientConfig{
		GRPCDialOptions: []grpc.DialOption{
			grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
				return net.Dial("tcp", fmt.Sprintf("%s:11001", host))
			}),
		},
		Reattach: &plugin.ReattachConfig{
			Protocol: "grpc",
			Addr:     addr,
		},
		HandshakeConfig: sp.Handshake,
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC,
		},
		Plugins: map[string]plugin.Plugin{
			"selflowRunner": &selflowRunnerProto.SelflowRunnerPlugin{},
		},
		Logger: hclog.Default(),
	})

	rpcClient, err := pluginClient.Client()
	if err != nil {
		return err
	}

	rawInitiator, err := rpcClient.Dispense("selflowRunner")
	if err != nil {
		return err
	}

	initiator, ok := rawInitiator.(selflowRunnerProto.SelflowRunner)
	if !ok {
		return errors.New("invalid protocol for selflow runner")
	}

	err = initiator.InitRunner(ctx, &containerSpawner{runId})
	pluginClient.Kill()

	return err
}
