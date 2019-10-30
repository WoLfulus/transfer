package commands

import (
	"github.com/foomo/htpasswd"
	"github.com/spf13/cobra"
	"github.com/wolfulus/transfer/transfer/config"
	"github.com/wolfulus/transfer/transfer/service"
)

func buildUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Manage additional user accounts",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(
		buildUserUpdate(),
		//buildUserRemove(),
	)
	return cmd
}

type userUpdateOptions struct {
	username string
	password string
}

func buildUserUpdate() *cobra.Command {
	var options userUpdateOptions

	cmd := &cobra.Command{
		Use:   "update <username> <password>",
		Short: "Updates or inserts a user",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			options.username = args[0]
			options.password = args[1]
			return runUserUpdate(options)
		},
	}

	return cmd
}

func runUserUpdate(options userUpdateOptions) error {
	svc, err := service.Get()
	if err != nil {
		return err
	}

	if config.IsService() {
		err := htpasswd.SetPassword(config.GetPasswordsFile(), options.username, options.password, htpasswd.HashBCrypt)
		if err != nil {
			return err
		}
		service.Log("User updated")
		return nil
	}

	return svc.Execute("/bin/transfer", "user", "update", options.username, options.password)
}
