package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/irth/edm/config"
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
			_ = cli

			f, err := os.Open("./config.json")
			if err != nil {
				return errors.Wrap(err, "opening config")
			}
			defer f.Close()
			cfg, err := config.Load(f, "hello")
			if err != nil {
				return errors.Wrap(err, "parsing config")
			}

			json.NewEncoder(os.Stdout).Encode(cfg)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
