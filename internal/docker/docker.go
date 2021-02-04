package docker

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// ContainerDetail contains information about a container
type ContainerDetail struct {
	ID        string
	Name      string
	Image     string
	IPAddress string
}

// CreateContainer creates a container from a given image
func CreateContainer(cli *client.Client, image string) (string, error) {
	ctx := context.Background()

	out, err := cli.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		return "", err
	}
	io.Copy(os.Stdout, out)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: image,
	}, nil, nil, nil, "")
	if err != nil {
		return "", err
	}

	return resp.ID, nil
}

// StartContainer starts an already created docker container
func StartContainer(cli *client.Client, containerID string) (types.ContainerJSON, error) {
	ctx := context.Background()
	if err := cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{}); err != nil {
		return types.ContainerJSON{}, err
	}
	container, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		return types.ContainerJSON{}, err
	}
	return container, nil
}

// GetContainerDetails returns the details of a container
func GetContainerDetails(cli *client.Client, container *types.ContainerJSON) ContainerDetail {
	return ContainerDetail{
		ID:        container.ID,
		Name:      container.Name,
		Image:     container.Image,
		IPAddress: container.NetworkSettings.IPAddress,
	}
}
