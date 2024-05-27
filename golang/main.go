package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "text",
				Aliases: []string{"t"},
				Usage:   "Create a text based on a prompt",
				Action: func(cCtx *cli.Context) error {
					CreateText()
					return nil
				},
			},
			{
				Name:    "image",
				Aliases: []string{"i"},
				Usage:   "Create an image based on a prompt",
				Action: func(cCtx *cli.Context) error {
					CreateImage()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
