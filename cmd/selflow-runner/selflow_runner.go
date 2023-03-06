package main

import (
	"github.com/hashicorp/go-plugin"
	selflowRunnerProto "github.com/selflow/selflow/internal/selflow-runner-proto"
	selflowPlugin "github.com/selflow/selflow/pkg/selflow-plugin"
	"google.golang.org/grpc"
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

		Listener: listener,
	})

}
