package workflow

import (
	"context"
	"fmt"
)

type ErrorStep struct {
	id     string
	status Status
}

func (s ErrorStep) GetOutput() map[string]string {
	return map[string]string{}
}

func (s *ErrorStep) Cancel() error {
	s.status = CANCELLED
	return nil
}

func (s ErrorStep) GetStatus() Status {
	return s.status
}

func (s ErrorStep) GetId() string {
	return s.id
}

func (s *ErrorStep) Execute(_ context.Context) (map[string]string, error) {
	s.status = ERROR

	return map[string]string{}, nil
}

func makeErrorStep(id string) (*ErrorStep, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("id must not be empty")
	}
	return &ErrorStep{id, CREATED}, nil
}
