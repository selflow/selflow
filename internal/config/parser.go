package config

import (
	"gopkg.in/yaml.v3"
)

func parse(config []byte) (*Flow, error) {
	validate := InitValidation()
	dst := TemplateFlow{}

	err := yaml.Unmarshal(config, &dst)
	if err != nil {
		return nil, err
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
