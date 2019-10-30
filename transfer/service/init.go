package service

import "github.com/docker/cli/cli/command"

var (
	cli   command.Cli
	debug bool
)

// Initialize the service
func Initialize(dockerCli command.Cli, dockerCliDebug bool) {
	cli = dockerCli
	debug = dockerCliDebug
}

// GetCLI gets the CLI instance
func GetCLI() command.Cli {
	return cli
}
