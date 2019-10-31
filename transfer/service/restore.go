package service

import (
	"context"

	cliImage "github.com/docker/cli/cli/command/image"
	"github.com/docker/docker/api/types"
)

// Restore an image from a repository
func Restore(registryImage string, localImage string) error {
	client := cli.Client()
	background := context.Background()
	cmd := cliImage.NewPullCommand(cli)

	err := cmd.RunE(cmd, []string{registryImage})
	if err != nil {
		return err
	}

	defer (func() {
		client.ImageRemove(background, registryImage, types.ImageRemoveOptions{
			Force: true,
		})
	})()

	return client.ImageTag(background, registryImage, localImage)
}
