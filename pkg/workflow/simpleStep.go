package workflow

import (
	"context"
	"fmt"
	"sync"
)

type SimpleStep struct {
	id       string
	status   Status
	statusMu sync.Mutex
}

func (s *SimpleStep) GetOutput() map[string]string {
	return map[string]string{}
}
func (s *SimpleStep) setStatus(status SimpleStatus) {
	s.statusMu.Lock()
	defer s.statusMu.Unlock()

	s.status = status
}

func (s *SimpleStep) Cancel() error {
	s.setStatus(CANCELLED)
	return nil
}

func (s *SimpleStep) GetStatus() Status {
	s.statusMu.Lock()
	defer s.statusMu.Unlock()
	return s.status
}

func (s *SimpleStep) GetId() string {
	return s.id
}

func (s *SimpleStep) Execute(_ context.Context) (map[string]string, error) {
	s.setStatus(SUCCESS)

	return map[string]string{}, nil
}

func makeSimpleStep(id string) (*SimpleStep, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("id must not be empty")
	}
	return newSimpleStep(id, CREATED), nil
}

func newSimpleStep(id string, status Status) *SimpleStep {
	return &SimpleStep{id: id, status: status}
}

var _ Step = &SimpleStep{}
