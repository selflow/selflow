package config

import (
  "encoding/json"
  "gopkg.in/yaml.v3"
)

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
