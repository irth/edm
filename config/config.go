package config

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/irth/edm/docker"
	"github.com/pkg/errors"
)

type container struct {
	Image  string `json:"image"`
	Mounts map[string]string
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

		for container, host := range spec.Mounts {
			co.Mounts = append(co.Mounts, docker.BindMount{
				Host:      host,
				Container: container,
			})
		}

		config.Containers = append(config.Containers, co)
	}

	return &config, nil
}
