package workflow

import (
	"context"
	"fmt"
)

type SimpleStep struct {
	id     string
	status Status
}

func (s *SimpleStep) GetOutput() map[string]string {
	return map[string]string{}
}

func (s *SimpleStep) Cancel() error {
	s.status = CANCELLED
	return nil
}

func (s *SimpleStep) GetStatus() Status {
	return s.status
}

func (s *SimpleStep) GetId() string {
	return s.id
}

func (s *SimpleStep) Execute(_ context.Context) (map[string]string, error) {
	s.status = SUCCESS

	return map[string]string{}, nil
}

func makeSimpleStep(id string) (*SimpleStep, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("id must not be empty")
	}
	return &SimpleStep{id, CREATED}, nil
}
