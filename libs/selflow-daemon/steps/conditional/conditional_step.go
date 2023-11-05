package conditional

import (
	"bytes"
	"context"
	"github.com/selflow/selflow/libs/core/workflow"
	"text/template"
)

type Step struct {
	workflow.Step
	condition string
}

var NoInputs = []string{
	"",
	"no",
	"false",
	"0",
}

func isTruthy(seq string) bool {
	for _, noInput := range NoInputs {
		if seq == noInput {
			return false
		}
	}

	return true
}

func NewConditionalStep(wrappedStep workflow.Step, condition string) workflow.Step {
	return &Step{
		Step:      wrappedStep,
		condition: condition,
	}
}

type TemplateValues struct {
	Needs map[string]map[string]string
}

func (step *Step) Execute(ctx context.Context) (map[string]string, error) {

	needs := ctx.Value(workflow.StepOutputContextKey).(map[string]map[string]string)

	tpl, err := template.New("").Parse(step.condition)
	if err != nil {
		return nil, err
	}

	buff := bytes.Buffer{}

	err = tpl.Execute(&buff, TemplateValues{needs})
	if err != nil {
		return nil, err
	}

	if isTruthy(buff.String()) {
		return step.Step.Execute(ctx)
	}

	step.SetStatus(workflow.CANCELLED)

	return map[string]string{}, nil
}
