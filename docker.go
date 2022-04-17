package main

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
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

func (c *DockerClient) Container(opts ContainerOptions) Container {
	return Container{
		ContainerOptions: opts,

		cli: c,
	}
}

func (c *DockerClient) pullImage(ctx context.Context, image string) error {
	reader, err := c.cli.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer reader.Close()
	io.Copy(os.Stdout, reader)
	return nil
}

type Mount interface {
	Mount() mount.Mount
}

type BindMount struct {
	Host      string
	Container string
}

func (b BindMount) Mount() mount.Mount {
	return mount.Mount{
		Type:   mount.TypeBind,
		Source: b.Host,
		Target: b.Container,
	}
}

func (c *DockerClient) createContainer(ctx context.Context, image string, name string, mounts []Mount) (string, error) {
	hostConfig := container.HostConfig{
		Mounts: make([]mount.Mount, len(mounts)),
	}
	for i, m := range mounts {
		hostConfig.Mounts[i] = m.Mount()
	}

	containerConfig := container.Config{
		Image: image,
	}

	resp, err := c.cli.ContainerCreate(ctx, &containerConfig, &hostConfig, nil, nil, name)
	if err != nil {
		return "", err
	}

	return resp.ID, nil
}

func (c *DockerClient) startContainer(ctx context.Context, containerID string) error {
	startOptions := types.ContainerStartOptions{}
	return c.cli.ContainerStart(ctx, containerID, startOptions)
}

func (c *DockerClient) stopContainer(ctx context.Context, containerID string, timeout *time.Duration) error {
	return c.cli.ContainerStop(ctx, containerID, timeout)
}

func (c *DockerClient) removeContainer(ctx context.Context, containerID string) error {
	removeOptions := types.ContainerRemoveOptions{}
	return c.cli.ContainerRemove(ctx, containerID, removeOptions)
}
