package selflow_plugin

import (
	"context"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type Architect interface {
	ValidateStepConfigSchema(context.Context, *ValidateStepConfigSchema_Request) (*ValidateStepConfigSchema_Response, error)
	RunStep(context.Context, *RunStep_Request) (*RunStep_Response, error)
}

type ArchitectPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	Impl Architect
}

func (a ArchitectPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	RegisterArchitectServer(server, &GRPCArchitectServer{
		Impl:   a.Impl,
		broker: broker,
	})
	return nil
}

func (a ArchitectPlugin) GRPCClient(_ context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &GRPCArchitectClient{
		broker:          broker,
		architectClient: NewArchitectClient(conn),
	}, nil
}

// CLIENT

type GRPCArchitectClient struct {
	broker          *plugin.GRPCBroker
	architectClient ArchitectClient
}

func (g *GRPCArchitectClient) ValidateStepConfigSchema(ctx context.Context, request *ValidateStepConfigSchema_Request) (*ValidateStepConfigSchema_Response, error) {
	return g.architectClient.ValidateStepConfigSchema(ctx, request)
}

func (g *GRPCArchitectClient) RunStep(ctx context.Context, request *RunStep_Request) (*RunStep_Response, error) {
	return g.architectClient.RunStep(ctx, request)
}

// SERVER

type GRPCArchitectServer struct {
	Impl   Architect
	broker *plugin.GRPCBroker
	UnimplementedArchitectServer
}

func (g *GRPCArchitectServer) ValidateStepConfigSchema(ctx context.Context, request *ValidateStepConfigSchema_Request) (*ValidateStepConfigSchema_Response, error) {
	return g.Impl.ValidateStepConfigSchema(ctx, request)
}

func (g *GRPCArchitectServer) RunStep(ctx context.Context, request *RunStep_Request) (*RunStep_Response, error) {
	return g.Impl.RunStep(ctx, request)
}

var _ plugin.GRPCPlugin = &ArchitectPlugin{}
var _ Architect = &GRPCArchitectClient{}
