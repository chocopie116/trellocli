package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", log.Ldate+log.Ltime+log.Lshortfile)

var Commands = []cli.Command{
	cli.Command{
		Name: "list",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "board",
				Usage: "Target board",
			},
		},
		Action: list,
	},
}

func main() {
	app := cli.NewApp()
	app.Commands = Commands
	app.Run(os.Args)
}

func list(c *cli.Context) error {
	b := c.String("board")
	logger.Print(b)

	return nil
}

