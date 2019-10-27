package commands

import (
	"os"

	"github.com/docker/cli/cli-plugins/plugin"
	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
	"github.com/wolfulus/transfer/transfer/service"
)

func rootCmd(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   name,
		Short: "Manage image transfers to other hosts",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(0)
		},
	}
	cmd.AddCommand(
		buildInstall(),
		buildStatus(),
		buildVersion(),
		buildUser(),
	)
	return cmd
}

// NewStandalone returns a new command configured to be used in standalone mode
func NewStandalone(dockerCli command.Cli) *cobra.Command {
	service.Initialize(dockerCli, os.Getenv("DOCKER_TRANSFER_DEBUG") != "")
	return rootCmd(os.Args[0])
}

// NewPlugin returns a new command configured to be used as a plugin
func NewPlugin(dockerCli command.Cli) *cobra.Command {
	service.Initialize(dockerCli, os.Getenv("DOCKER_TRANSFER_DEBUG") != "")
	cmd := rootCmd("transfer")
	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		return plugin.PersistentPreRunE(cmd, args)
	}
	return cmd
}
