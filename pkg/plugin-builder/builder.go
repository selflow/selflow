package plugin_builder

import (
	"fmt"
	"github.com/hashicorp/go-plugin"
	selflowPlugin "github.com/selflow/selflow/pkg/selflow-plugin"
	"log"
	"net"
	"os"
)

const PortEnv = "SELFLOW_PORT"

func assertEnvironmentVariable(variableName string) string {
	value, ok := os.LookupEnv(variableName)
	if !ok || value == "" {
		panic(fmt.Sprintf("Environment variable [%s] must be defined", variableName))
	}
	return value
}

func checkEnvironmentConfig() {
	assertEnvironmentVariable(selflowPlugin.Handshake.MagicCookieKey)
	assertEnvironmentVariable(PortEnv)
}

type ServePluginConfig struct {
	ArchitectPlugin *selflowPlugin.ArchitectPlugin
	BasicPlugin     *selflowPlugin.BasicPlugin
}

func buildPluginMap(config ServePluginConfig) map[string]plugin.Plugin {
	plugins := make(map[string]plugin.Plugin)

	if config.ArchitectPlugin != nil {
		plugins["architect"] = config.ArchitectPlugin
	}

	if config.BasicPlugin != nil {
		plugins["basicPlugin"] = config.BasicPlugin
	}

	return plugins
}

func ServePlugin(config ServePluginConfig) {
	checkEnvironmentConfig()

	port := os.Getenv(PortEnv)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	plugins := buildPluginMap(config)

	log.Printf("Start serving plugin on port %s with options %v\n", port)

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: selflowPlugin.Handshake,
		Plugins:         plugins,
		GRPCServer:      plugin.DefaultGRPCServer,
		Listener:        listener,
	})

}
