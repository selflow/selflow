package workflow

import (
	"context"
	"errors"
)

type errorStep struct {
	*SimpleStep
}

func (s *errorStep) GetOutput() map[string]string {
	return map[string]string{}
}

func (s *errorStep) Execute(_ context.Context) (map[string]string, error) {
	return map[string]string{}, errors.New("some-error")
}

func makeErrorStep(id string) Step {
	return &errorStep{newSimpleStep(id, CREATED)}
}
