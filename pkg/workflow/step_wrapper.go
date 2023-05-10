package workflow

import (
	"context"
)

type stepWrapper struct {
	Step
}

func (s *stepWrapper) Execute(ctx context.Context) (map[string]string, error) {
	s.SetStatus(RUNNING)
	outputs, err := s.Step.Execute(ctx)
	if !s.GetStatus().IsFinished() {
		if err != nil {
			s.SetStatus(ERROR)
			return outputs, err
		}

		s.SetStatus(SUCCESS)
	}

	return outputs, nil
}

func wrapStep(step Step) Step {
	return &stepWrapper{step}
}

var _ Step = &SimpleStep{}
