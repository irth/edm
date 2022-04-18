package docker

import (
	"github.com/docker/docker/api/types/mount"
)

type Mount interface {
	Prepare() error
	Mount() mount.Mount
}

type BindMount struct {
	Host      string
	Container string
}

func (b BindMount) Prepare() error {
	return nil
}

func (b BindMount) Mount() mount.Mount {
	return mount.Mount{
		Type:   mount.TypeBind,
		Source: b.Host,
		Target: b.Container,
	}
}
