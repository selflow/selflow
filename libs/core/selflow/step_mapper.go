package selflow

import (
	"github.com/selflow/selflow/libs/core/workflow"
	"github.com/selflow/selflow/libs/selflow-daemon/config"
)

type StepMapper interface {
	MapStep(stepId string, definition config.StepDefinition) (workflow.Step, error)
}
