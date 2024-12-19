package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "webgate"
	app.Version = AppVersion
	app.Description = "Webgate"
	app.Usage = "Route different clients to different services defined in your configuration file"

	app.Authors = []cli.Author{
		{
			Name:  "Mike Savochkin (crioto)",
			Email: "mike@crioto.com",
		},
	}

	app.Copyright = "2024 (c) crioto.com. All Right Reserved"

	app.Commands = []cli.Command{
		{
			Name:  "serve",
			Usage: "Start service",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "config",
					Usage:       "Path to config file",
					Value:       "/etc/noerrorcode.yaml",
					Destination: &ConfigFilePath,
				},
			},
			Action: RunService,
		},
	}

	app.Run(os.Args)
}
