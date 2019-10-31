package commands

import (
	"github.com/spf13/cobra"
	"github.com/wolfulus/transfer/transfer/service"
)

type restoreOptions struct {
	username      string
	password      string
	registry      string
	registryImage string
	localImage    string
}

func buildRestore() *cobra.Command {
	options := restoreOptions{}

	cmd := &cobra.Command{
		Use:   "restore <remote> <local>",
		Args:  cobra.MinimumNArgs(2),
		Short: "Restores a remote image locally",
		RunE: func(cmd *cobra.Command, args []string) error {
			options.registryImage = args[0]
			options.localImage = args[1]
			return runRestore(options)
		},
		Hidden: true,
	}

	flags := cmd.Flags()

	flags.StringVar(&options.registry, "registry", "", "Registry hostname")
	flags.StringVar(&options.username, "username", "", "Auth username")
	flags.StringVar(&options.password, "password", "", "Auth password")

	return cmd
}

func runRestore(options restoreOptions) error {
	err := service.Authenticate(options.registry, options.username, options.password)
	if err != nil {
		service.Error("failed to store auth information for %s", options.registry)
	}

	err = service.Restore(options.registryImage, options.localImage)
	if err != nil {
		service.Error("failed to restore image '%s' as '%s'", options.registryImage, options.localImage)
		return err
	}
	return nil
}
