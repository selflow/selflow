package localexec

import (
	"context"
	"fmt"
	"github.com/selflow/selflow/libs/core/sflog"
	"github.com/selflow/selflow/libs/core/workflow"
	"log/slog"
	"os/exec"
)

type Step struct {
	workflow.SimpleStep
	config Config
}

type Config struct {
	Command          string
	Env              map[string]string
	WorkingDirectory string
}

func (step *Step) Execute(ctx context.Context) (map[string]string, error) {
	step.SetStatus(workflow.RUNNING)

	envs := make([]string, len(step.config.Env))
	for envKey, envValue := range step.config.Env {
		envs = append(envs, fmt.Sprintf("%s=%s", envKey, envValue))
	}

	cmd := exec.Command("bash", "-c", step.config.Command)

	// flows
	sout := sflog.WriterFromLogger(sflog.GetLoggerFromContext(ctx), slog.LevelDebug, slog.String("command", cmd.Path), slog.Any("args", cmd.Args), slog.String("channel", "stdout"))
	serr := sflog.WriterFromLogger(sflog.GetLoggerFromContext(ctx), slog.LevelDebug, slog.String("command", cmd.Path), slog.Any("args", cmd.Args), slog.String("channel", "stderr"))

	cmd.Stdout = sout
	cmd.Stderr = serr

	slog.DebugContext(ctx, "Start command", "command", cmd.String())
	err := cmd.Run()
	if err != nil {
		slog.ErrorContext(ctx, "Command failed", "error", err)
		return nil, err
	}

	return nil, nil
}

func (step *Step) GetOutput() map[string]string {
	return map[string]string{}
}
