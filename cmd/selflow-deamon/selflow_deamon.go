package main

import (
	"context"
	"github.com/selflow/selflow/internal/config"
	"github.com/selflow/selflow/pkg/runners"
	"os"
)

func main() {

	configAsBytes, err := os.ReadFile("/etc/selflow/config.yaml")
	if err != nil {
		panic(err)
	}

	flow, err := config.Parse(configAsBytes)
	if err != nil {
		panic(err)
	}

	runners.StartRunner(context.Background(), *flow)

}
