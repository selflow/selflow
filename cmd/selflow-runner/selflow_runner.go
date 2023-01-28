package main

import (
	"context"
	"github.com/hashicorp/go-plugin"
	selflowPlugin "github.com/selflow/selflow/pkg/selflow-plugin"
	"log"
	"net"
)

func contains[T comparable](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}

	return false
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:7001")
	if err != nil {
		panic(err)
	}

	//pluginClient, err := plugin.NewRPCClient(conn, selflowPlugin.PluginMap)
	pluginClient := plugin.NewClient(&plugin.ClientConfig{
		//GRPCDialOptions: []grpc.DialOption{
		//	grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
		//		return net.Dial("tcp", "127.0.0.1:12345")
		//	}),
		//},
		Reattach: &plugin.ReattachConfig{
			Protocol: "grpc",
			Addr:     addr,
		},
		HandshakeConfig: selflowPlugin.Handshake,
		Plugins:         selflowPlugin.PluginMap,
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC,
		},
	})

	rpcClient, err := pluginClient.Client()
	if err != nil {
		panic(err)
	}

	err = rpcClient.Ping()
	if err != nil {
		panic(err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("basicPlugin")
	if err != nil {
		panic(err)
	}

	pluginDef := raw.(selflowPlugin.Plugin)
	_, err = pluginDef.GetPluginSchema()
	if err != nil {
		panic(err)
	}

	pluginSchemaResponse, err := pluginDef.GetPluginSchema()

	log.Printf("%v", pluginSchemaResponse.PluginTypes)

	if contains(pluginSchemaResponse.PluginTypes, selflowPlugin.PluginType_ARCHITECT) {
		log.Printf("Plugin is an architect")

		raw, err := rpcClient.Dispense("architect")
		if err != nil {
			panic(err)
		}

		architectPlugin := raw.(selflowPlugin.Architect)

		stepConfig := []byte("{\"name\": \"Tonton Anthony\"}")

		ctx := context.TODO()

		rep, err := architectPlugin.ValidateStepConfigSchema(ctx, &selflowPlugin.ValidateStepConfigSchema_Request{
			StepConfig: stepConfig,
		})

		if !rep.Valid {
			log.Printf("fail to validate config : %v", rep.Diagnotics)
			return
		}

		architectPlugin.RunStep(ctx, &selflowPlugin.RunStep_Request{
			StepConfig: stepConfig,
		})
	}

}
