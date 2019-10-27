package service

import (
	"context"
	"errors"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/wolfulus/transfer/transfer/version"
)

var (
	// ErrServiceNotFound when service could not be found
	ErrServiceNotFound = errors.New("service not found")

	// ErrMultipleServicesRunning when multiple services are running
	ErrMultipleServicesRunning = errors.New("multiple services running")
)

// Service represents a service running service instance
type Service struct {
	Container types.Container
	Version   string
	Status    string
	Outdated  bool
	Managed   bool
	Running   bool
}

// Get the service instance
func Get() (Service, error) {
	service := Service{
		Status: StatusUnknown,
	}

	client := cli.Client()
	background := context.Background()

	filters := filters.NewArgs()
	filters.Add("label", LabelVersion)

	containers, err := client.ContainerList(background, types.ContainerListOptions{
		All:     true,
		Filters: filters,
	})

	if err != nil {
		return service, err
	}

	if len(containers) == 0 {
		service.Status = StatusNotFound
		return service, ErrServiceNotFound
	} else if len(containers) > 1 {
		service.Status = StatusMultiple
		return service, ErrMultipleServicesRunning
	}

	container := containers[0]

	_, managed := container.Labels[LabelManaged]
	service.Managed = managed
	service.Version = container.Labels[LabelVersion]
	service.Outdated = container.Labels[LabelVersion] != version.Version
	service.Running = container.State == "running" || container.State == "restarting" || container.State == "paused"
	service.Container = container
	service.Status = StatusOK
	return service, nil
}

// Uninstall the service instance
func (svc *Service) Uninstall() {
	client := cli.Client()
	background := context.Background()

	timeout := time.Second * 15

	stopped := false

	Debug("Uninstalling container %v", svc.Container.ID)

	err := client.ContainerStop(background, svc.Container.ID, &timeout)
	if err != nil {
		Debug("Error stopping container: %v", err)
	} else {
		done, errc := client.ContainerWait(background, svc.Container.ID, container.WaitConditionNotRunning)
		select {
		case err := <-errc:
			if err != nil {
				Debug("Error waiting for stop: %v", err)
			}
		case <-done:
			Debug("Container stopped")
			stopped = true
		}
	}

	if stopped == false {
		Debug("Killing container")
		err = client.ContainerKill(background, svc.Container.ID, "SIGKILL")
		if err != nil {
			Debug("Error killing container: %v", err)
		} else {
			done, errc := client.ContainerWait(background, svc.Container.ID, container.WaitConditionNotRunning)
			select {
			case err := <-errc:
				if err != nil {
					Log("Error waiting for stop: %v", err)
				}
			case <-done:
				Log("Container killed")
			}
		}
	}

	err = client.ContainerRemove(background, svc.Container.ID, types.ContainerRemoveOptions{
		Force: true,
	})

	if err != nil {
		Debug("Error removing container: %v", err)
	} else {
		done, errc := client.ContainerWait(background, svc.Container.ID, container.WaitConditionRemoved)
		select {
		case err := <-errc:
			if err != nil {
				Debug("Error waiting for removal: %v", err)
			}
		case <-done:
			Debug("Container removed")
		}
	}
}
