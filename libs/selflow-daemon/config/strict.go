package config

type Metadata map[string]interface{}

type Flow struct {
	Metadata Metadata `validate:"required"`
	Inputs   Inputs   `validate:"required,identifierKey"`
	Outputs  Outputs  `validate:"required,identifierKey"`
	Plugins  Plugins  `validate:"required,identifierKey"`
	Workflow Workflow `validate:"required"`
}

type Input struct {
	Type        string `validate:"required"`
	Description string `validate:"required"`
	Required    bool   `validate:"required"`
	Secret      bool   `validate:"required"`
	Default     string `validate:"required"`
}
type Inputs map[string]Input

type Outputs map[string]string

type PluginConfiguration struct {
	Version string                 `validate:"required"`
	Config  map[string]interface{} `validate:"required"`
}

type Plugins map[string]PluginConfiguration

type PluginStepConfig map[string]interface{}

type With PluginStepConfig

type StepDefinition struct {
	If            string   `validate:"required,template"`
	Timeout       string   `validate:"required,duration"`
	Matrix        []string `validate:"required,dive,len=1"`
	Needs         []string `validate:"required"`
	Kind          string   `validate:"required"`
	OnErrorIgnore bool     `validate:"required"`
	With          With     `validate:"required"`
}

type Steps map[string]StepDefinition

type Workflow struct {
	Timeout string `validate:"required,duration"`
	Steps   Steps  `validate:"required,identifierKey"`
}
