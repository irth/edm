package docker

import (
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
