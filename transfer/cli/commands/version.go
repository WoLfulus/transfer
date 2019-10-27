package commands

import (
	"github.com/wolfulus/transfer/transfer/service"
	"github.com/wolfulus/transfer/transfer/version"

	"github.com/spf13/cobra"
)

type versionOptions struct {
}

func buildVersion() *cobra.Command {

	options := versionOptions{}

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show the transfer plugin version",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runVersion(options)
		},
	}

	//flags := cmd.Flags()
	//flags.StringVar(&options.variable, "name", "default", "Some variable")

	return cmd
}

func runVersion(options versionOptions) error {
	service.Log("Transfer version: %s", version.Version)
	return nil
}
