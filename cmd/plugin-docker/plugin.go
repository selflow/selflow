package main

import (
	pluginBuilder "github.com/selflow/selflow/pkg/plugin-builder"
	sp "github.com/selflow/selflow/pkg/selflow-plugin"
	"log"
)

type Summary struct {
}

func (s Summary) GetPluginSchema() (*sp.GetPluginSchema_Response, error) {
	return &sp.GetPluginSchema_Response{
		PluginTypes: []sp.PluginType{
			sp.PluginType_ARCHITECT,
		},
		StepSchemas: map[string]*sp.Schema{
			"image": {
				Type:         "string",
				NestedGroups: nil,
				Description:  "Docker image to use",
				Required:     true,
				Deprecated:   false,
			},
			"commands": {
				Type:         "string",
				NestedGroups: nil,
				Description:  "Commands to run on the container",
				Required:     true,
				Deprecated:   false,
			},
			"environment": {
				Type:         "map",
				NestedGroups: nil,
				Description:  "Map of environment variables",
				Required:     true,
				Deprecated:   false,
			},
		},
	}, nil
}

func (s Summary) ValidatePluginConfigSchema(config []byte) (*sp.ValidatePluginConfigSchema_Response, error) {
	log.Print(string(config))

	return &sp.ValidatePluginConfigSchema_Response{
		Valid:      true,
		Diagnotics: nil,
	}, nil
}

func main() {
	pluginBuilder.ServePlugin(pluginBuilder.ServePluginConfig{
		BasicPlugin:     &sp.BasicPlugin{Impl: &Summary{}},
		ArchitectPlugin: &sp.ArchitectPlugin{Impl: &DockerArchitect{}},
	})
}

var _ sp.Plugin = &Summary{}
