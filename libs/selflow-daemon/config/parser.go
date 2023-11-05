package config

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
)

// Parse try to parse a metadata file from a binary json. If it fails try to parse
// it from yaml and return the Flow or an error if it fails
func Parse(config []byte) (*Flow, error) {
	validate := InitValidation()
	dst := TemplateFlow{}

	err := json.Unmarshal(config, &dst)
	if err != nil {
		err = yaml.Unmarshal(config, &dst)
		if err != nil {
			return nil, err
		}
	}

	strictFlow, err := dst.ToStrictConfig()
	if err != nil {
		return nil, err
	}

	err = validate.Struct(strictFlow)
	if err != nil {
		return nil, err
	}

	return strictFlow, nil
}
