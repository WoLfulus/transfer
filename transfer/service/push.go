package service

import (
	"context"
	"fmt"

	cliImage "github.com/docker/cli/cli/command/image"
	"github.com/docker/docker/api/types"
	"github.com/wolfulus/transfer/transfer/tags"
)

type pushOptions struct {
	remote    string
	untrusted bool
}

// Push a image to a repository
func Push(image string, server string) error {
	client := cli.Client()
	context := context.Background()

	// Check if image exists
	_, _, err := client.ImageInspectWithRaw(context, image)
	if err != nil {
		return err
	}

	// Generate the target tag with original name encoded
	tag, err := tags.Encode(image)
	if err != nil {
		return err
	}

	// Tag the image to another image
	remoteImage := fmt.Sprintf("%s/wolfulus/transferred:%s", server, tag)
	err = client.ImageTag(context, image, remoteImage)
	if err != nil {
		return err
	}

	// Make sure we remove it later
	defer (func() {
		client.ImageRemove(context, remoteImage, types.ImageRemoveOptions{})
	})()

	/*
		closer, err := client.ImagePush(context, remoteImage, types.ImagePushOptions{
			All:          true,
			RegistryAuth: "{}",
		})
		if err != nil {
			return err
		}
	*/
	cmd := cliImage.NewPushCommand(cli)

	return cmd.RunE(cmd, []string{remoteImage})
}
