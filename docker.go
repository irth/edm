package main

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
)

type DockerClient struct {
	cli *client.Client
}

func NewDockerClient() (*DockerClient, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	return &DockerClient{cli: cli}, nil
}

func (c *DockerClient) PullImage(ctx context.Context, image string) error {
	reader, err := c.cli.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		return errors.Wrap(err, "pulling image")
	}
	defer reader.Close()
	io.Copy(os.Stdout, reader)
	return nil
}

func (c *DockerClient) CreateContainer(ctx context.Context, image string) (string, error) {
	hostConfig := container.HostConfig{}
	containerConfig := container.Config{
		Image: image,
	}

	resp, err := c.cli.ContainerCreate(ctx, &containerConfig, &hostConfig, nil, nil, "")
	if err != nil {
		return "", errors.Wrap(err, "creating container")
	}

	return resp.ID, nil
}

func (c *DockerClient) StartContainer(ctx context.Context, containerID string) error {
	startOptions := types.ContainerStartOptions{}
	err := c.cli.ContainerStart(ctx, containerID, startOptions)
	if err != nil {
		return errors.Wrap(err, "starting container")
	}

	return nil
}
