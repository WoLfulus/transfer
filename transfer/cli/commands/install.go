package commands

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/wolfulus/transfer/transfer/service"
)

type installOptions struct {
	username string
	password string
	force    bool
}

func buildInstall() *cobra.Command {

	options := installOptions{}

	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install transfer server in the current host",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runInstall(options)
		},
	}

	flags := cmd.Flags()

	flags.StringVarP(&options.username, "username", "u", "", "The initial user username")
	cobra.MarkFlagRequired(flags, "username")

	flags.StringVarP(&options.password, "password", "p", "", "The initial user password")
	cobra.MarkFlagRequired(flags, "password")

	flags.BoolVarP(&options.force, "force", "f", false, "Forces the installation")

	return cmd
}

func runInstall(options installOptions) error {

	svc, err := service.Get()
	if err != nil {
		return err
	}

	if svc.Status == service.StatusUnknown {
		service.Log("Service status could not be determined. Aborting.")
		os.Exit(1)
		return nil
	}

	if svc.Status == service.StatusMultiple {
		service.Log("Impossible to install because multiple services are already running.")
		service.Log("Please remove all previous service containers first.")
		os.Exit(1)
		return nil
	}

	if svc.Status == service.StatusOK {
		if options.force == false {
			service.Log("There is another service installed. Please uninstall the previous service first or use --force to override the current service.")
			os.Exit(1)
			return nil
		}

		if svc.Managed == false {
			service.Log("The current service cannot be replaced because it wasn't created by transfer CLI.")
			os.Exit(1)
			return nil
		}

		svc.Uninstall()
	}

	//service.Install(username, password)
	return nil
}
