package selflow_runner_proto

import (
	"context"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
	"log"
	"net"
)

type SelflowRunner interface {
	InitRunner(ctx context.Context, spawner ContainerSpawner) error
}

// GRPCSelflowRunnerServer Server
type GRPCSelflowRunnerServer struct {
	Impl   SelflowRunner
	broker *plugin.GRPCBroker
	UnimplementedSelflowRunnerServer
}

func (s *GRPCSelflowRunnerServer) InitRunner(ctx context.Context, req *InitRunner_Request) (*InitRunner_Response, error) {
	conn, err := s.broker.Dial(req.ContainerSpawnServer)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := &GRPCContainerSpawnerClient{NewContainerSpawnerClient(conn)}
	return &InitRunner_Response{}, s.Impl.InitRunner(ctx, client)
}

// GRPCSelflowRunnerClient Client
type GRPCSelflowRunnerClient struct {
	broker *plugin.GRPCBroker
	client SelflowRunnerClient
}

func (c *GRPCSelflowRunnerClient) InitRunner(ctx context.Context, spawner ContainerSpawner) error {
	containerSpawnerServer := &GRPCContainerSpawnerServer{Impl: spawner}
	var server *grpc.Server

	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		server = grpc.NewServer(opts...)
		RegisterContainerSpawnerServer(server, containerSpawnerServer)
		return server
	}

	listener, err := net.Listen("tcp", "selflow-daemon:11002")
	log.Printf("Start lisenting on port :11002\n")
	if err != nil {
		return err
	}

	brokerID := c.broker.NextId()
	go func() {
		listener, err = c.broker.AcceptWithCustomListener(brokerID, listener)
		if err != nil {
			log.Printf("[ERR] plugin: plugin acceptAndServe error: %s", err)
			return
		}
		c.broker.AcceptAndServeWithCustomListener(brokerID, serverFunc, listener)
	}()

	_, err = c.client.InitRunner(ctx, &InitRunner_Request{
		ContainerSpawnServer: brokerID,
	})

	return err
}

// SelflowRunnerPlugin Plugin
type SelflowRunnerPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	Impl SelflowRunner
}

func (p SelflowRunnerPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	RegisterSelflowRunnerServer(server, &GRPCSelflowRunnerServer{
		Impl:   p.Impl,
		broker: broker,
	})
	return nil
}

func (p SelflowRunnerPlugin) GRPCClient(_ context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {

	return &GRPCSelflowRunnerClient{
		broker: broker,
		client: NewSelflowRunnerClient(conn),
	}, nil
}
