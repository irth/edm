package docker

import (
	"context"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/pkg/errors"
)

func (c Container) Down(ctx context.Context) error {
	err := c.cli.ContainerStop(ctx, c.Name, nil)
	if err != nil {
		if strings.Contains(err.Error(), "No such container:") {
			return nil
		}
		return err
	}

	err = c.cli.ContainerRemove(ctx, c.Name, types.ContainerRemoveOptions{})
	if err != nil {
		return err
	}

	for _, m := range c.Mounts {
		err = m.Dispose()
		if err != nil {
			return errors.Wrapf(err, "disposing mount %+v", m)
		}
	}
	return nil
}
