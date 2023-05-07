package main

import (
	"fmt"
	"github.com/selflow/selflow/cmd/selflow-daemon/server"
	"github.com/selflow/selflow/cmd/selflow-daemon/server/proto"
	"github.com/selflow/selflow/internal/sfenvironment"
	"github.com/selflow/selflow/pkg/logger/systemfile"
	"github.com/selflow/selflow/pkg/sflog"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"path"
)

func setupLogger() {
	logger := sflog.LoggerFromEnv("selflow-daemon")
	sflog.SetDefaultLogger(logger)
}

func main() {
	setupLogger()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", sfenvironment.GetDaemonPort()))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	proto.RegisterDaemonServer(s, &server.Server{
		LogFactory: systemfile.NewLogFactory(path.Join(os.Getenv("PWD"), "tmp")),
	})

	log.Printf("[INFO] Start listening at %v\n", lis.Addr())
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
