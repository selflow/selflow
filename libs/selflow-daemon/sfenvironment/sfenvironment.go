package sfenvironment

import (
	"github.com/selflow/selflow/libs/selflow-daemon/envutils"
	"os"
	"strings"
)

const (
	DaemonPortEnvKey          = "SELFLOW_DAEMON_PORT"
	DaemonPortEnvDefaultValue = "10011"

	DaemonDebugPortEnvKey          = "SELFLOW_DAEMON_DEBUG_PORT"
	DaemonDebugPortEnvDefaultValue = ""

	DaemonNameEnvKey          = "SELFLOW_DAEMON_NAME"
	DaemonNameEnvDefaultValue = "selflow-daemon"

	DaemonBaseDirectoryEnvKey          = "SELFLOW_DAEMON_BASE_DIRECTORY"
	DaemonBaseDirectoryEnvDefaultValue = "/etc/selflow"

	DaemonNetworkEnvKey = "SELFLOW_DAEMON_NETWORK"

	DaemonImageEnvKey          = "SELFLOW_DAEMON_IMAGE"
	DaemonImageEnvDefaultValue = "selflow-daemon:latest"

	DaemonHostBaseDirectoryEnvKey = "SELFLOW_DAEMON_HOST_BASED_DIRECTORY"

	UseJsonLogEnvKey        = "JSON_LOGS"
	UseJsonLogsDefaultValue = "FALSE"

	LogLevelEnvKey       = "LOG_LEVEL"
	LogLevelDefaultValue = "INFO"
)

var (
	UseJsonLogs = strings.ToUpper(envutils.GetEnv(UseJsonLogEnvKey, UseJsonLogsDefaultValue)) == "TRUE"
	LogLevel    = envutils.GetEnv(LogLevelEnvKey, LogLevelDefaultValue)
)

func EnvOrDefault(key string, defaultValue string) string {
	environmentValue := os.Getenv(key)
	if environmentValue == "" {
		return defaultValue
	}
	return environmentValue
}

func GetDaemonPort() string {
	return EnvOrDefault(DaemonPortEnvKey, DaemonPortEnvDefaultValue)
}

func GetDaemonDebugPort() string {
	return EnvOrDefault(DaemonDebugPortEnvKey, DaemonDebugPortEnvDefaultValue)
}

func GetDaemonName() string {
	return EnvOrDefault(DaemonNameEnvKey, DaemonNameEnvDefaultValue)
}

func GetDaemonBaseDirectory() string {
	return EnvOrDefault(DaemonBaseDirectoryEnvKey, DaemonBaseDirectoryEnvDefaultValue)
}

func GetDaemonHostBaseDirectory() string {
	return EnvOrDefault(DaemonHostBaseDirectoryEnvKey, GetDaemonBaseDirectory())
}

func GetDaemonNetwork() string {
	return EnvOrDefault(DaemonNetworkEnvKey, GetDaemonName())
}

func GetDaemonImage() string {
	return EnvOrDefault(DaemonImageEnvKey, DaemonImageEnvDefaultValue)
}
