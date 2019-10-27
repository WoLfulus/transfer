package commands

import (
	"github.com/spf13/cobra"
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
		Use:   "update",
		Short: "Updates or inserts a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runUserUpdate(options)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&options.username, "username", "", "The target username")
	flags.StringVar(&options.password, "password", "", "The new user's password ")

	cmd.MarkFlagRequired("username")

	return cmd
}

func runUserUpdate(options userUpdateOptions) error {
	//
	service.Log("Update user %s to use password %s", options.username, options.password)
	return nil
}
