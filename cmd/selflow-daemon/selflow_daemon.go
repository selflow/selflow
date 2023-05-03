package main

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/selflow/selflow/cmd/selflow-daemon/server"
	"github.com/selflow/selflow/cmd/selflow-daemon/server/proto"
	"github.com/selflow/selflow/internal/sfenvironment"
	"google.golang.org/grpc"
	"log"
	"net"
)

func setupLogger() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:            "selflow-daemon",
		Output:          nil,
		JSONFormat:      false,
		IncludeLocation: false,
		TimeFormat:      "2006-01-02 15:04:05",
		Color:           hclog.ForceColor,
		Level:           hclog.Debug,
	})

	hclog.SetDefault(logger)

	log.SetOutput(logger.StandardWriter(&hclog.StandardLoggerOptions{InferLevels: true}))
	log.SetPrefix("")
	log.SetFlags(0)
}

func main() {
	setupLogger()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", sfenvironment.GetDaemonPort()))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	proto.RegisterDaemonServer(s, &server.Server{})

	log.Printf("[INFO] Start listening at %v\n", lis.Addr())
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
