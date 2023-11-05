package main

import (
	"fmt"
	"github.com/selflow/selflow/apps/selflow-daemon/server"
	"github.com/selflow/selflow/apps/selflow-daemon/server/proto"
	"github.com/selflow/selflow/libs/core/sflog"
	"github.com/selflow/selflow/libs/selflow-daemon/logger/systemfile"
	"github.com/selflow/selflow/libs/selflow-daemon/runPersistence/sqlite"
	"github.com/selflow/selflow/libs/selflow-daemon/sfenvironment"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"path"
)

func init() {
	sflog.Init(slog.String("process", "selflow-daemon"))
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", sfenvironment.GetDaemonPort()))
	if err != nil {
		panic(err)
	}

	slog.Debug("Connecting to SQLITE database", "database", "history.db")
	runPersistence, err := sqlite.NewSqliteRunPersistence(path.Join(sfenvironment.GetDaemonBaseDirectory(), "history.db"))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	proto.RegisterDaemonServer(s, &server.Server{
		LogFactory:     systemfile.NewLogFactory(path.Join(sfenvironment.GetDaemonBaseDirectory(), "tmp")),
		RunPersistence: runPersistence,
	})

	slog.Info("Start GRPC server", "address", lis.Addr())
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
