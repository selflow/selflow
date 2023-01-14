package main

import (
	"encoding/json"
	"fmt"
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
			"name": {
				Type:         "string",
				NestedGroups: nil,
				Description:  "Name of the person to say hello",
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

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloArchitect struct {
}

func (h *HelloArchitect) ValidateStepConfigSchema(request *sp.ValidateStepConfigSchema_Request) (*sp.ValidateStepConfigSchema_Response, error) {
	helloRequest := HelloRequest{}

	err := json.Unmarshal(request.GetStepConfig(), &helloRequest)
	if err != nil {
		return &sp.ValidateStepConfigSchema_Response{
			Valid: false,
			Diagnotics: []*sp.Diagnostic{
				{
					Type:    sp.DiagnosticType_ERROR,
					Message: fmt.Sprintf("fail to load config : %v", err),
				},
			},
		}, nil
	}

	return &sp.ValidateStepConfigSchema_Response{
		Valid:      true,
		Diagnotics: []*sp.Diagnostic{},
	}, nil
}

func (h *HelloArchitect) RunStep(request *sp.RunStep_Request) (*sp.RunStep_Response, error) {

	helloRequest := HelloRequest{}

	err := json.Unmarshal(request.GetStepConfig(), &helloRequest)
	if err != nil {
		panic(err)
	}

	log.Printf("Hello %v", helloRequest.Name)

	return &sp.RunStep_Response{}, nil
}

func main() {
	pluginBuilder.ServePlugin(pluginBuilder.ServePluginConfig{
		BasicPlugin:     &sp.BasicPlugin{Impl: &Summary{}},
		ArchitectPlugin: &sp.ArchitectPlugin{Impl: &HelloArchitect{}},
	})
}

var _ sp.Plugin = &Summary{}
var _ sp.Architect = &HelloArchitect{}
