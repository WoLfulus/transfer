package service

import (
	"github.com/docker/cli/cli/command/container"
)

// Execute a command inside the service image
func (svc *Service) Execute(command string, commandArgs ...string) error {
	args := append([]string{"-it", svc.Container.ID, command}, commandArgs...)
	cmd := container.NewExecCommand(cli)
	cmd.SetArgs(args)
	_, err := cmd.ExecuteC()
	return err
}
