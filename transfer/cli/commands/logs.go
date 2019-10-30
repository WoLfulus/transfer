package commands

import (
	"fmt"

	container "github.com/docker/cli/cli/command/container"
	"github.com/spf13/cobra"
	"github.com/wolfulus/transfer/transfer/service"
)

type logOptions struct {
	follow bool
	tail   string
}

func buildLog() *cobra.Command {

	options := logOptions{}

	cmd := &cobra.Command{
		Use:   "logs",
		Short: "Shows service logs",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runLog(options)
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&options.follow, "follow", "f", false, "Follow log output")
	flags.StringVar(&options.tail, "tail", "all", "Number of lines to show from the end of the logs")

	return cmd
}

func runLog(options logOptions) error {
	cli := service.GetCLI()
	svc, err := service.Get()
	if err != nil {
		return err
	}

	args := []string{}
	if options.follow {
		args = append(args, "--follow")
	}

	if options.tail != "" {
		args = append(args, fmt.Sprintf("--tail=%s", options.tail))
	}

	args = append(args, svc.Container.ID)

	cmd := container.NewLogsCommand(cli)
	cmd.SetArgs(args)
	cmd.SetOutput(cli.Out())
	_, err = cmd.ExecuteC()
	return err
}
