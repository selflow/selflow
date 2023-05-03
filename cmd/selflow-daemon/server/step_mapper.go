package server

import (
	"github.com/selflow/selflow/internal/config"
	"github.com/selflow/selflow/pkg/workflow"
)

type StepMapper interface {
	MapStep(stepId string, definition config.StepDefinition) (workflow.Step, error)
}
