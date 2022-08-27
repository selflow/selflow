package workflow

import (
	"context"
)

type Step interface {
	GetStatus() Status
	Execute(context context.Context) (map[string]string, error)
	Cancel() error
	GetId() string
	GetOutput() map[string]string
}

type StepWithInit interface {
	Step
	Init(context context.Context) error
}

type StepWithTearDown interface {
	Step
	TearDown(context context.Context) error
}

// Workflow contains a list of steps and behave like any other step
type Workflow interface {
	Step
	AddStep(step *Step, dependencies []*Step) error
}

func sliceContainsStep(slice []Step, element Step) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}
