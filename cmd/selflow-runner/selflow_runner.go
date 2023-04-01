package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	selflowRunnerProto "github.com/selflow/selflow/internal/selflow-runner-proto"
	dockerStep "github.com/selflow/selflow/pkg/docker-step"
	selflowPlugin "github.com/selflow/selflow/pkg/selflow-plugin"
	"github.com/selflow/selflow/pkg/wfbuilder"
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

	selflowRunnerPlugin := &SelflowRunnerPlugin{
		workflowBuilder: wfbuilder.Builder{
			StepBuilderMap: map[string]wfbuilder.StepMapper{
				"docker": dockerStep.NewDockerStep,
			},
		},
		configFileLocation: "/etc/selflow/config.json",
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: selflowPlugin.Handshake,
		Plugins: map[string]plugin.Plugin{
			"runner": &selflowRunnerProto.SelflowRunnerPlugin{Impl: selflowRunnerPlugin},
		},
		GRPCServer: func(options []grpc.ServerOption) *grpc.Server {
			return plugin.DefaultGRPCServer(options)
		},
		Logger: hclog.Default(),

		Listener: listener,
	})
}
