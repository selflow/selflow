package main

import (
	"context"
	"github.com/hashicorp/go-hclog"
	"github.com/selflow/selflow/internal/config"
	"github.com/selflow/selflow/pkg/runners"
	"log"
	"os"
)

func setupLogger() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:            "selflow-daemon",
		Output:          nil,
		JSONFormat:      false,
		IncludeLocation: false,
		TimeFormat:      "2006-01-02 15:04:05",
		Color:           hclog.ForceColor,
		Level:           hclog.Debug,
	})

	hclog.SetDefault(logger)

	log.SetOutput(logger.StandardWriter(&hclog.StandardLoggerOptions{InferLevels: true}))
	log.SetPrefix("")
	log.SetFlags(0)
}

func main() {
	setupLogger()

	configAsBytes, err := os.ReadFile("/etc/selflow/config.yaml")
	if err != nil {
		panic(err)
	}

	flow, err := config.Parse(configAsBytes)
	if err != nil {
		panic(err)
	}

	ctx := context.WithValue(context.Background(), "debug-port", "40001")

	runners.StartRunner(ctx, *flow)

}
