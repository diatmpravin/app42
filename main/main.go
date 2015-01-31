package main

import (
	"github.com/codegangsta/cli"
	"github.com/diatmpravin/cli/app42/commands"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "app42"
	app.Usage = "A command line tool to interact with App42PaaS"
	app.Version = "1.0.0.alpha"
	app.Author = "Pravin Mishra"
	app.Email = "pravinmishra88@gmail.com"

	app.Commands = []cli.Command{
		{
			Name:  "apps",
			Usage: "List all the deployed applications with their meta details",
			Action: func(c *cli.Context) {
				cmd := commands.NewApps()
				cmd.Run(c)
			},
		},
	}

	app.Run(os.Args)

}
