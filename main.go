package main

import (
	"log"
	"os"
	"time"

	"github.com/irth/edm/docker"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "edm",
		Usage: "idk",
		Action: func(ctx *cli.Context) error {
			log.Println("hello")

			cli, err := docker.NewDockerClient()
			if err != nil {
				return errors.Wrap(err, "creating docker client")
			}

			container := cli.Container(docker.ContainerOptions{
				Name:  "test_container_1",
				Image: "nginx:latest",
				Mounts: []docker.Mount{
					docker.BindMount{"/tmp/a", "/tmp/b"},
				},
			})

			err = container.Up(ctx.Context)
			if err != nil {
				return err
			}

			<-time.After(2 * time.Second)

			err = container.Down(ctx.Context)
			if err != nil {
				return err
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
