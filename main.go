package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func runContainer() error {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return errors.Wrap(err, "creating docker client")
	}

	// PARAMS
	ctx := context.Background()
	image := "nginx:latest"

	// 1. Pull Image
	reader, err := cli.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		return errors.Wrap(err, "pulling image")
	}
	defer reader.Close()
	io.Copy(os.Stdout, reader)

	// 2. Create Container
	hostConfig := container.HostConfig{}
	containerConfig := container.Config{
		Image: image,
	}

	resp, err := cli.ContainerCreate(ctx, &containerConfig, &hostConfig, nil, nil, "")
	if err != nil {
		return errors.Wrap(err, "creating container")
	}

	containerID := resp.ID

	// 3. Start Container
	startOptions := types.ContainerStartOptions{}
	err = cli.ContainerStart(ctx, containerID, startOptions)
	if err != nil {
		return errors.Wrap(err, "starting container")
	}

	log.Printf("started container %s", containerID)
	return nil
}

func main() {
	app := &cli.App{
		Name:  "edm",
		Usage: "idk",
		Action: func(ctx *cli.Context) error {
			log.Println("hello")
			return runContainer()
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
