package docker

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/pkg/errors"
)

func (c Container) Up(ctx context.Context) error {
	err := c.Down(ctx)
	if err != nil {
		return err
	}

	reader, err := c.cli.ImagePull(ctx, c.Image, types.ImagePullOptions{})
	if err != nil {
		return errors.Wrapf(err, "error while pulling %s", c.Image)
	}

	defer reader.Close()
	io.Copy(os.Stdout, reader)

	hostConfig := container.HostConfig{
		Mounts: make([]mount.Mount, len(c.Mounts)),
	}
	for i, m := range c.Mounts {
		err = m.Prepare()
		mount := m.Mount()
		if err != nil {
			return errors.Wrapf(err, "error while preparing mount %s", mount.Target)
		}
		hostConfig.Mounts[i] = mount
	}

	containerConfig := container.Config{
		Image: c.Image,
	}

	resp, err := c.cli.ContainerCreate(
		ctx,
		&containerConfig, &hostConfig, nil,
		nil,
		c.Name)
	if err != nil {
		return errors.Wrapf(err, "couldn't create container %s", c.Name)
	}

	return c.cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
}
