package main

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
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
		return err
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
		return "", err
	}

	return resp.ID, nil
}

func (c *DockerClient) StartContainer(ctx context.Context, containerID string) error {
	startOptions := types.ContainerStartOptions{}
	return c.cli.ContainerStart(ctx, containerID, startOptions)
}

func (c *DockerClient) StopContainer(ctx context.Context, containerID string, timeout *time.Duration) error {
	return c.cli.ContainerStop(ctx, containerID, timeout)
}

func (c *DockerClient) RemoveContainer(ctx context.Context, containerID string) error {
	removeOptions := types.ContainerRemoveOptions{}
	return c.cli.ContainerRemove(ctx, containerID, removeOptions)
}
