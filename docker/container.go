package docker

import "github.com/docker/docker/client"

type ContainerOptions struct {
	Name   string
	Image  string
	Mounts []Mount
}

type Container struct {
	ContainerOptions

	cli *client.Client
}
