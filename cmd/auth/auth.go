package auth

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func AuthCommand() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "register",
				Aliases: []string{"r"},
				Usage:   "Register to the system",
				Action:  RegisterActions,
			},
			{
				Name:    "login",
				Aliases: []string{"l"},
				Usage:   "Login to the system",
				Action:  LoginActions,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
