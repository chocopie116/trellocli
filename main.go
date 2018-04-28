package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/adlio/trello"
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
			cli.StringFlag{
				Name:  "app_key",
				Usage: "Trello API app key",
			},
			cli.StringFlag{
				Name:  "token",
				Usage: "Trello API token",
			},
		},
		Action: list,
	},
}



func main() {
	app := cli.NewApp()
	app.Commands = Commands
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "app_key",
		},
	}
	app.Run(os.Args)
}

func list(c *cli.Context) error {
	k := c.String("app_key")
	t := c.String("token")

	client := trello.NewClient(k, t)

	b := c.String("board")
	bc, _ := client.GetBoard(b, trello.Defaults())

	logger.Print(bc)

	return nil
}

