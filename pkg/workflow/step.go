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

func sliceContainsStep(slice []Step, element Step) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}
