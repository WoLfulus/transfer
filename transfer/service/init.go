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
