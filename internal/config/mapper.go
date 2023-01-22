package config

import (
	"encoding/json"
	"fmt"
)

func (f TemplateFlow) extractMetadata() (Metadata, error) {
	var metadata map[string]interface{}

	templateAsJson, err := json.Marshal(f)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(templateAsJson, &metadata)
	if err != nil {
		return nil, err
	}
	delete(metadata, "Inputs")
	delete(metadata, "Default")
	delete(metadata, "Workflow")
	delete(metadata, "Plugins")
	delete(metadata, "Outputs")
	return metadata, nil
}

func (f TemplateFlow) compileInputs() (Inputs, error) {
	inputs := Inputs{}
	if f.Inputs == nil {
		return inputs, nil
	}

	for k, v := range f.Inputs {
		inputs[k] = Input{
			Type:        v.Type,
			Description: v.Description,
			Required:    v.Required,
			Secret:      v.Secret,
			Default:     v.Default,
		}
	}
	return inputs, nil
}

func (f TemplateFlow) compileOutputs() (Outputs, error) {
	outputs := map[string]string{}
	if f.Outputs == nil {
		return outputs, nil
	}

	for k, v := range f.Outputs {
		outputs[k] = v
	}
	return outputs, nil
}

func (f TemplateFlow) compilePlugins() (Plugins, error) {
	plugins := Plugins{}
	if f.Plugins == nil {
		return plugins, nil
	}

	for k, v := range f.Plugins {
		plugins[k] = PluginConfiguration{
			Version: v.Version,
			Config:  v.Config,
		}
	}
	return plugins, nil
}

func mapToWith(src interface{}, w *With) error {
	optionsAsJson, err := json.Marshal(src)
	if err != nil {
		return fmt.Errorf("fail to convert to json : %v", err)
	}
	err = json.Unmarshal(optionsAsJson, &w)
	if err != nil {
		return fmt.Errorf("fail to process json : %v", err)
	}
	return nil
}

func (t TemplateStepDefinition) compileStep(stepsDefaultValues TemplateDefault) (*StepDefinition, error) {
	defaultOptions := With{}
	options := With{}

	if stepsDefaultValues != nil {
		defaultValues, founded := stepsDefaultValues[t.Kind]
		if founded {
			err := mapToWith(defaultValues, &defaultOptions)
			if err != nil {
				return nil, fmt.Errorf("fail to process %v default options : %v", t.Kind, err)
			}
		}
	}

	err := mapToWith(t.With, &options)
	if err != nil {
		return nil, fmt.Errorf("fail to process options : %v", err)
	}

	with := defaultOptions.merge(options)

	return &StepDefinition{
		If:            t.If,
		Timeout:       t.Timeout,
		Matrix:        t.Matrix,
		Needs:         t.Needs,
		Kind:          t.Kind,
		OnErrorIgnore: t.OnErrorIgnore,
		With:          with,
	}, nil
}

func (w With) merge(w2 With) With {
	for k, v := range w2 {
		w[k] = v
	}
	return w
}

func (f TemplateFlow) compileWorkflow() (*Workflow, error) {
	steps := Steps{}

	for stepName, templateStep := range f.Workflow.Steps {
		step, err := templateStep.compileStep(f.Default)
		if err != nil {
			return nil, fmt.Errorf("fail to process step [%v] : %v", stepName, err)
		}
		steps[stepName] = *step
	}

	return &Workflow{
		Timeout: f.Workflow.Timeout,
		Steps:   steps,
	}, nil
}

func (f TemplateFlow) ToStrictConfig() (*Flow, error) {
	metadata, err := f.extractMetadata()
	if err != nil {
		return nil, err
	}

	inputs, err := f.compileInputs()
	if err != nil {
		return nil, err
	}

	outputs, err := f.compileOutputs()
	if err != nil {
		return nil, err
	}

	plugins, err := f.compilePlugins()
	if err != nil {
		return nil, err
	}

	workflow, err := f.compileWorkflow()
	if err != nil {
		return nil, err
	}

	return &Flow{
		Metadata: metadata,
		Inputs:   inputs,
		Outputs:  outputs,
		Plugins:  plugins,
		Workflow: *workflow,
	}, nil
}
