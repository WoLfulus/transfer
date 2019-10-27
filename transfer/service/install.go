package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/wolfulus/transfer/transfer/version"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

// Install the service into the host machine
func Install(username string, password string, port int, address string, env []string) error {

	client := cli.Client()
	background := context.Background()

	hostConfig := container.HostConfig{}

	if port > 0 {
		hostBinding := nat.PortBinding{
			HostIP:   address,
			HostPort: strconv.Itoa(port),
		}
		containerPort, err := nat.NewPort("tcp", "5000")
		if err != nil {
			return nil
		}
		hostConfig.PortBindings = nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}
	}

	config := container.Config{
		Image: version.FQDN,
		Labels: map[string]string{
			LabelVersion: version.Version,
			LabelManaged: "yes",
		},
		Env: append([]string{
			fmt.Sprintf("MANAGEMENT_USERNAME=%s", username),
			fmt.Sprintf("MANAGEMENT_PASSWORD=%s", password),
		}, env...),
	}

	cont, err := client.ContainerCreate(background, &config, &hostConfig, nil, "")
	if err != nil {
		Debug("Failed to create container: %s", err)
		return err
	}

	err = client.ContainerStart(background, cont.ID, types.ContainerStartOptions{})
	if err != nil {
		Debug("Failed to start container: %s", err)
		return err
	}

	return nil
}
