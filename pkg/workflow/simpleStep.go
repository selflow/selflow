package workflow

import (
	"context"
	"fmt"
	"sync"
)

type SimpleStep struct {
	Id       string
	Status   Status
	statusMu sync.Mutex
}

func (s *SimpleStep) GetOutput() map[string]string {
	return map[string]string{}
}
func (s *SimpleStep) SetStatus(status Status) {
	s.statusMu.Lock()
	defer s.statusMu.Unlock()

	s.Status = status
}

func (s *SimpleStep) Cancel() error {
	s.SetStatus(CANCELLED)
	return nil
}

func (s *SimpleStep) GetStatus() Status {
	s.statusMu.Lock()
	defer s.statusMu.Unlock()
	return s.Status
}

func (s *SimpleStep) GetId() string {
	return s.Id
}

func (s *SimpleStep) Execute(_ context.Context) (map[string]string, error) {
	s.SetStatus(SUCCESS)

	return map[string]string{}, nil
}

func makeSimpleStep(id string) (*SimpleStep, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("id must not be empty")
	}
	return newSimpleStep(id, CREATED), nil
}

func newSimpleStep(id string, status Status) *SimpleStep {
	return &SimpleStep{Id: id, Status: status}
}

var _ Step = &SimpleStep{}
