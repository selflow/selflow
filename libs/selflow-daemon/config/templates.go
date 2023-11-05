package config

type TemplateFlow struct {
	Name        string           `yaml:"name"`
	Description string           `yaml:"description"`
	Author      string           `yaml:"author"`
	Inputs      TemplateInputs   `yaml:"inputs"`
	Outputs     TemplateOutputs  `yaml:"outputs"`
	Plugins     TemplatePlugins  `yaml:"plugins"`
	Default     TemplateDefault  `yaml:"default"`
	Workflow    TemplateWorkflow `yaml:"workflow"`
}

type TemplateInput struct {
	Type        string `yaml:"type"`
	Description string `yaml:"description"`
	Required    bool   `yaml:"required"`
	Secret      bool   `yaml:"secret"`
	Default     string `yaml:"default"`
}
type TemplateInputs map[string]TemplateInput

type TemplateOutputs map[string]string

type TemplatePluginVersionConfigDetailed struct {
	Version string                 `yaml:"version"`
	Config  map[string]interface{} `yaml:"config"`
}

type TemplatePluginVersionConfigVersionOnly string

type TemplatePluginVersionConfig TemplatePluginVersionConfigDetailed

type TemplatePlugins map[string]TemplatePluginVersionConfig

type TemplatePluginStepConfig map[string]interface{}

type TemplateDefault map[string]TemplatePluginStepConfig

type TemplateWith TemplatePluginStepConfig

type TemplateStepDefinition struct {
	If            string       `yaml:"if"`
	Timeout       string       `yaml:"timeout"`
	Matrix        []string     `yaml:"matrix"`
	Needs         []string     `yaml:"needs"`
	Kind          string       `yaml:"kind"`
	OnErrorIgnore bool         `yaml:"on-error-ignore"`
	With          TemplateWith `yaml:"with"`
}

type TemplateSteps map[string]TemplateStepDefinition

type TemplateWorkflow struct {
	Timeout string        `yaml:"timeout"`
	Steps   TemplateSteps `yaml:"steps"`
}
