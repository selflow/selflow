package runners

import (
	"context"
	"errors"
	"fmt"
	"github.com/eknkc/basex"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/selflow/selflow/internal/config"
	"log"
	"time"
)

const (
	runnerContainerBaseName     = "selflow-runner"
	daemonNetwork               = "selflow-daemon-network"
	allowedIdentifierCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var MaximumRetriesExceeded = errors.New("maximum retries exceeded")

func runners(name string) string {
	result := "runners " + name
	return result
}

func generateRunId() string {

	uu := uuid.New()

	encoding, err := basex.NewEncoding(allowedIdentifierCharacters)
	if err != nil {
		// The base must compile
		panic(err)
	}

	return encoding.Encode(uu[:])
}

func initContainerWithRetries(ctx context.Context, runId string) error {
	for remainingTries := 10; remainingTries > 0; remainingTries-- {
		err := initContainer(ctx, runId)
		if err == nil || !errors.Is(err, plugin.ErrProcessNotFound) {
			return err
		}

		time.Sleep(5 * time.Second)
	}
	return MaximumRetriesExceeded
}

// StartRunner aims to be run by the daemon to start and follow
// a runner and provide it needed methods
func StartRunner(ctx context.Context, flow config.Flow) {
	runId := generateRunId()

	logger := hclog.Default().Named(fmt.Sprintf("runner-%s", runId))

	runnerExitCode, err := startRunnerContainer(ctx, flow, runId, logger)

	if err != nil {
		logger.Error("fail to start runner", err)
		return
	}

	err = initContainerWithRetries(ctx, runId)
	if err != nil {
		log.Printf("[ERROR] fail to initialize runner %s : %v\n", runId, err)
	}

	<-runnerExitCode

}
