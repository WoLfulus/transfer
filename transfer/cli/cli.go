package cli

import (
	"fmt"
	"os"

	"github.com/wolfulus/transfer/transfer/cli/commands"
	"github.com/wolfulus/transfer/transfer/version"

	"github.com/docker/cli/cli-plugins/manager"
	"github.com/docker/cli/cli-plugins/plugin"
	"github.com/docker/cli/cli/command"
	cliflags "github.com/docker/cli/cli/flags"
	"github.com/spf13/cobra"
)

// Execute the command line application
func Execute() {

	isPlugin := false
	if os.Getenv("DOCKER_CLI_PLUGIN_ORIGINAL_CLI_COMMAND") != "" {
		isPlugin = true
	} else if len(os.Args) >= 2 && os.Args[1] == manager.MetadataSubcommandName {
		isPlugin = true
	}

	if isPlugin {
		plugin.Run(func(dockerCli command.Cli) *cobra.Command {
			return commands.NewPlugin(dockerCli)
		}, manager.Metadata{
			SchemaVersion: "0.1.0",
			Vendor:        "WoLfulus",
			Version:       version.Version,
			//Experimental:  os.Getenv("TRANSFER_EXPERIMENTAL") != "",
		})
		return
	}

	dockerCli, err := command.NewDockerCli()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	dockerCli.Initialize(cliflags.NewClientOptions())
	cmd := commands.NewStandalone(dockerCli)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
