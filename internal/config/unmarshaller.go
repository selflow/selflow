package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

func (p *TemplatePluginVersionConfig) UnmarshalYAML(value *yaml.Node) error {
	var versionOnly TemplatePluginVersionConfigVersionOnly
	if err := value.Decode(&versionOnly); err == nil {
		p.Version = string(versionOnly)
		p.Config = make(map[string]interface{})
		return nil
	}

	var fullPluginDefinition TemplatePluginVersionConfigDetailed
	if err := value.Decode(&fullPluginDefinition); err == nil {
		p.Version = fullPluginDefinition.Version
		p.Config = fullPluginDefinition.Config
		return nil
	}

	return fmt.Errorf("syntax err : invalid plugin version")
}
