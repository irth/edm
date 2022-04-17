package main

import (
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "edm",
		Usage: "idk",
		Action: func(ctx *cli.Context) error {
			log.Println("hello")

			cli, err := NewDockerClient()
			if err != nil {
				return errors.Wrap(err, "creating docker client")
			}

			container := cli.Container(ContainerOptions{
				Name:  "test_container_1",
				Image: "nginx:latest",
				Mounts: []Mount{
					BindMount{"/tmp/a", "/tmp/b"},
				},
			})

			err = container.Up(ctx.Context)
			if err != nil {
				return err
			}

			<-time.After(5 * time.Second)

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
