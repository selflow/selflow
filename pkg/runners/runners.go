package runners

import (
  "context"
  "encoding/hex"
  "fmt"
  "github.com/eknkc/basex"
  "github.com/google/uuid"
  "github.com/hashicorp/go-hclog"
  "github.com/selflow/selflow/internal/config"
  "strings"
  "time"
)

const (
  runnerContainerBaseName     = "selflow-runner"
  daemonNetwork               = "selflow-daemon-network"
  allowedIdentifierCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func runners(name string) string {
  result := "runners " + name
  return result
}

func generateRunId() string {

  uuidWithoutHyphens := strings.Replace(uuid.New().String(), "-", "", -1)

  uuidBytes, err := hex.DecodeString(uuidWithoutHyphens)
  if err != nil {
    // The uuid must parse
    panic(err)
  }

  encoding, err := basex.NewEncoding(allowedIdentifierCharacters)
  if err != nil {
    // The base must compile
    panic(err)
  }

  return encoding.Encode(uuidBytes)
}

// StartRunner aims to be run by the daemon to start and follow
// a runner and provide it needed methods
func StartRunner(ctx context.Context, flow config.Flow) {
  runId := generateRunId()

  logger := hclog.New(&hclog.LoggerOptions{
    Name: fmt.Sprintf("selflow-runner-%s", runId),
  })

  runnerExitCode, err := startRunnerContainer(ctx, flow, runId, logger)

  time.Sleep(3 * time.Second)

  if err != nil {
    logger.Error("fail to start runner", err)
    return
  }

  err = initContainer(ctx, runId)
  if err != nil {
    logger.Error("fail to initialize runner", err)
    return
  }

  <-runnerExitCode

}
