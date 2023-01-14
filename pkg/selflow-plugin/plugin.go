package selflow_plugin

import (
	"context"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type Plugin interface {
	GetPluginSchema() (*GetPluginSchema_Response, error)
	ValidatePluginConfigSchema(config []byte) (*ValidatePluginConfigSchema_Response, error)
}

type BasicPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	Impl Plugin
}

func (a BasicPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	RegisterPluginServer(server, &GRPCPluginServer{
		Impl:   a.Impl,
		broker: broker,
	})
	return nil
}

func (a BasicPlugin) GRPCClient(_ context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &GRPCPluginClient{
		broker:       broker,
		pluginClient: NewPluginClient(conn),
	}, nil
}

// CLIENT

type GRPCPluginClient struct {
	pluginClient PluginClient
	broker       *plugin.GRPCBroker
}

func (g *GRPCPluginClient) GetPluginSchema() (*GetPluginSchema_Response, error) {
	res, err := g.pluginClient.GetPluginSchema(context.Background(), &GetPluginSchema_Request{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (g *GRPCPluginClient) ValidatePluginConfigSchema(config []byte) (*ValidatePluginConfigSchema_Response, error) {
	res, err := g.pluginClient.ValidatePluginConfigSchema(context.Background(), &ValidatePluginConfigSchema_Request{
		PluginConfig: config,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SERVER

type GRPCPluginServer struct {
	Impl   Plugin
	broker *plugin.GRPCBroker
	UnimplementedPluginServer
}

func (s *GRPCPluginServer) GetPluginSchema(_ context.Context, _ *GetPluginSchema_Request) (*GetPluginSchema_Response, error) {
	v, err := s.Impl.GetPluginSchema()
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (s *GRPCPluginServer) ValidatePluginConfigSchema(_ context.Context, request *ValidatePluginConfigSchema_Request) (*ValidatePluginConfigSchema_Response, error) {
	v, err := s.Impl.ValidatePluginConfigSchema(request.PluginConfig)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (s *GRPCPluginServer) mustEmbedUnimplementedPluginServer() {
}
