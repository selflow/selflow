package container

import (
	"bytes"
	"text/template"
)

type TemplateValues struct {
	Needs map[string]map[string]string
}

func withTemplate(templateString string, values map[string]map[string]string) (string, error) {
	tpl, err := template.New("").Parse(templateString)
	if err != nil {
		return "", err
	}

	buff := bytes.Buffer{}

	err = tpl.Execute(&buff, TemplateValues{values})
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
