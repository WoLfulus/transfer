package commands

import (
	"github.com/spf13/cobra"
	"github.com/wolfulus/transfer/transfer/config"
	"github.com/wolfulus/transfer/transfer/service"
)

type imageOptions struct {
	image  string
	target string
	as     string
}

func buildImage() *cobra.Command {

	options := imageOptions{}

	cmd := &cobra.Command{
		Use:   "image <server> <image name>",
		Args:  cobra.MinimumNArgs(2),
		Short: "Sends an image to the specified server",
		RunE: func(cmd *cobra.Command, args []string) error {
			options.target = args[0]
			options.image = args[1]
			return runImage(options)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&options.as, "as", "", "Rename image on target machine")

	return cmd
}

func runImage(options imageOptions) error {
	image := options.image
	alias := options.image
	if options.as != "" {
		alias = options.as
	}

	server := config.GetServerFromAlias(options.target)
	err := service.Push(image, alias, server)
	if err != nil {
		service.Error("Failed to push image '%s' to '%s'", options.image, server)
		return err
	}
	return nil
}
