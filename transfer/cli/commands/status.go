package commands

import (
	"github.com/spf13/cobra"
	"github.com/wolfulus/transfer/transfer/service"
)

type statusOptions struct {
}

func buildStatus() *cobra.Command {

	options := installOptions{}

	cmd := &cobra.Command{
		Use:   "status",
		Short: "Shows the status of transfer server in this host",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runStatus(options)
		},
	}

	//flags := cmd.Flags()
	//flags.StringVar(&options.variable, "name", "default", "Some variable")

	return cmd
}

func runStatus(options installOptions) error {
	svc, err := service.Get()
	if err != nil {
		return err
	}

	if svc.Status == service.StatusNotFound {
		service.Log("Status: service not found.")
	} else if svc.Status == service.StatusUnknown {
		service.Log("Status: unknown")
	} else if svc.Status == service.StatusMultiple {
		service.Log("Status: multiple instances running")
	} else {
		if svc.Running {
			service.Log("Status: running")
		} else {
			service.Log("Status: stopped")
		}
	}

	return nil
}
