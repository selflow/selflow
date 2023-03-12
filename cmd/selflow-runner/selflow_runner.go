package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	selflowRunnerProto "github.com/selflow/selflow/internal/selflow-runner-proto"
	selflowPlugin "github.com/selflow/selflow/pkg/selflow-plugin"
	"google.golang.org/grpc"
	"log"
	"net"
)

func setupLogger() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:            "",
		Output:          nil,
		JSONFormat:      false,
		IncludeLocation: false,
		TimeFormat:      "",
		DisableTime:     true,
		Color:           hclog.AutoColor,
	})

	hclog.SetDefault(logger)

	log.SetPrefix("")
	log.SetFlags(0)
}

func main() {
	setupLogger()

	listener, err := net.Listen("tcp", ":11001")
	if err != nil {
		panic(err)
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: selflowPlugin.Handshake,
		Plugins: map[string]plugin.Plugin{
			"runner": &selflowRunnerProto.SelflowRunnerPlugin{Impl: &SelflowRunnerPlugin{}},
		},
		GRPCServer: func(options []grpc.ServerOption) *grpc.Server {
			return plugin.DefaultGRPCServer(options)
		},
		Logger: hclog.Default(),

		Listener: listener,
	})

	log.Println("terminated")

}
