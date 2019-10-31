package service

import (
	reg "github.com/docker/cli/cli/command/registry"
)

// Authenticate into a registry
func Authenticate(registry string, username string, password string) error {
	cmd := reg.NewLoginCommand(cli)
	cmd.SetOutput(cli.Out())
	cmd.SetArgs([]string{"--username", username, "--password", password, registry})
	_, err := cmd.ExecuteC()
	return err
}
