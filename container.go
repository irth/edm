package main

import (
	"context"
	"strings"
)

type ContainerOptions struct {
	Name   string
	Image  string
	Mounts []Mount
}

type Container struct {
	ContainerOptions

	cli *DockerClient
}

func (c Container) Up(ctx context.Context) error {
	err := c.Down(ctx)
	if err != nil {
		return err
	}

	err = c.cli.pullImage(ctx, c.Image)
	if err != nil {
		return err
	}

	id, err := c.cli.createContainer(ctx, c.Image, c.Name, c.Mounts)
	if err != nil {
		return err
	}

	return c.cli.startContainer(ctx, id)
}

func (c Container) Down(ctx context.Context) error {
	err := c.cli.stopContainer(ctx, c.Name, nil)
	if err != nil {
		if strings.Contains(err.Error(), "No such container:") {
			return nil
		}
		return err
	}

	return c.cli.removeContainer(ctx, c.Name)
}
