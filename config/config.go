package config

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/irth/edm/docker"
	"github.com/pkg/errors"
)

type container struct {
	Image  string           `json:"image"`
	Mounts map[string]mount `json:"mounts"`
}

type mount struct {
	Type   string `json:"type"`
	Source string `json:"source"`
}

type config struct {
	Containers map[string]container `json:"containers"`
}

type Config struct {
	Name       string
	Containers []docker.ContainerOptions
}

func Load(r io.Reader, name string) (*Config, error) {
	var parsed config
	err := json.NewDecoder(r).Decode(&parsed)
	if err != nil {
		return nil, errors.Wrap(err, "decoding json")
	}

	var config Config
	config.Name = name

	config.Containers = make([]docker.ContainerOptions, 0, len(parsed.Containers))
	for containerName, spec := range parsed.Containers {
		co := docker.ContainerOptions{
			Name:   fmt.Sprintf("edm_%s_%s", name, containerName),
			Image:  spec.Image,
			Mounts: make([]docker.Mount, 0, len(spec.Mounts)),
		}

		for container, mountSpec := range spec.Mounts {
			var mount docker.Mount
			switch mountSpec.Type {
			case "bind":
				mount = docker.BindMount{
					Host:      mountSpec.Source,
					Container: container,
				}
				// TODO: template mount
			default:
				return nil, fmt.Errorf("unknown mount type for %s: %s", container, mountSpec.Type)
			}
			co.Mounts = append(co.Mounts, mount)
		}

		config.Containers = append(config.Containers, co)
	}

	return &config, nil
}
