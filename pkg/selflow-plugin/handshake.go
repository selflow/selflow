package selflow_plugin

import "github.com/hashicorp/go-plugin"

var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "SELFLOW_PLUGIN",
	MagicCookieValue: "OWE5ZDBiNGEtOWE1My00ODI0LTkxYzAtNGM3NDRiNTQ2ZDJhCg==",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"architect":   &ArchitectPlugin{},
	"basicPlugin": &BasicPlugin{},
}
