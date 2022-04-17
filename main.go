package main

import (
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "edm",
		Usage: "idk",
		Action: func(ctx *cli.Context) error {
			log.Println("hello")
			image := "nginx:latest"

			cli, err := NewDockerClient()
			if err != nil {
				return errors.Wrap(err, "creating docker client")
			}

			err = cli.PullImage(ctx.Context, image)
			if err != nil {
				return err
			}

			id, err := cli.CreateContainer(ctx.Context, image)
			if err != nil {
				return err
			}

			err = cli.StartContainer(ctx.Context, id)
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
