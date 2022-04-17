package main

import (
	"context"

	"github.com/irth/edm/config"
	"github.com/irth/edm/docker"
	"github.com/pkg/errors"
)

func Up(ctx context.Context, cli *docker.DockerClient, cfg *config.Config) error {
	var err error
	for _, containerOptions := range cfg.Containers {
		c := cli.Container(containerOptions)
		err = c.Up(ctx)
		if err != nil {
			return errors.Wrapf(err, "while starting container %s", c.Name)
		}
	}

	return nil
}

func Down(ctx context.Context, cli *docker.DockerClient, cfg *config.Config) error {
	var err error
	for _, containerOptions := range cfg.Containers {
		c := cli.Container(containerOptions)
		err = c.Down(ctx)
		if err != nil {
			return errors.Wrapf(err, "while starting container %s", c.Name)
		}
	}

	return nil
}
