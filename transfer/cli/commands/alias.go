package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wolfulus/transfer/transfer/config"
)

type aliasOptions struct {
	alias  string
	server string
}

func buildAlias() *cobra.Command {
	options := aliasOptions{}
	cmd := &cobra.Command{
		Use:   "alias <name> [new server]",
		Args:  cobra.MinimumNArgs(1),
		Short: "Gets or sets an alias to a remote server",
		RunE: func(cmd *cobra.Command, args []string) error {
			options.alias = args[0]
			if len(args) > 1 {
				options.server = args[1]
			}
			return runAlias(options)
		},
	}
	return cmd
}

func runAlias(options aliasOptions) error {
	if options.server != "" {
		config.SetServerAlias(options.alias, options.server)
	}
	fmt.Printf("Alias '%s' is set to '%s'\n", options.alias, config.GetServerFromAlias(options.alias))
	return nil
}
