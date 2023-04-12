package workflow

import (
	"context"
)

type stepWrapper struct {
	Step
}

func (s *stepWrapper) Execute(ctx context.Context) (map[string]string, error) {
	outputs, err := s.Step.Execute(ctx)
	if err != nil {
		s.SetStatus(ERROR)
		return nil, err
	}

	if !s.GetStatus().IsFinished() {
		s.SetStatus(SUCCESS)
	}

	return outputs, nil
}

func wrapStep(step Step) Step {
	return &stepWrapper{step}
}

var _ Step = &SimpleStep{}
