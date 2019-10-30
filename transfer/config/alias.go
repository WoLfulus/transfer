package config

import (
	"fmt"

	"github.com/wolfulus/transfer/transfer/service"
	"github.com/wolfulus/transfer/transfer/version"
)

// GetServerFromAlias will get a server from its alias
func GetServerFromAlias(alias string) string {
	cli := service.GetCLI()
	value, ok := cli.ConfigFile().PluginConfig(version.Plugin, fmt.Sprintf("alias.%s", alias))
	if ok {
		return value
	}
	return alias
}

// SetServerAlias will set an alias to a server
func SetServerAlias(alias string, server string) {
	cli := service.GetCLI()
	cli.ConfigFile().SetPluginConfig(version.Plugin, fmt.Sprintf("alias.%s", alias), server)
	cli.ConfigFile().Save()
}
