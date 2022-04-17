package docker

import (
	"context"
	"strings"

	"github.com/docker/docker/api/types"
)

func (c Container) Down(ctx context.Context) error {
	err := c.cli.ContainerStop(ctx, c.Name, nil)
	if err != nil {
		if strings.Contains(err.Error(), "No such container:") {
			return nil
		}
		return err
	}

	return c.cli.ContainerRemove(ctx, c.Name, types.ContainerRemoveOptions{})
}
