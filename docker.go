package main

import (
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

		cli: c.cli,
	}
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
