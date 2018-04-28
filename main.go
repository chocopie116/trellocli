package main

import (
	"log"
	"os"

	"github.com/adlio/trello"
	"github.com/k0kubun/pp"
	"github.com/urfave/cli"
)

var logger = log.New(os.Stdout, "", log.Ldate+log.Ltime+log.Lshortfile)

var commands = []cli.Command{
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
	app.Commands = commands
	app.Flags = []cli.Flag{

		cli.StringFlag{
			Name: "app_key",
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}

	logger.Print("done")
	pp.Print("ok")
}

func list(c *cli.Context) error {
	k := c.String("app_key")
	t := c.String("token")

	client := trello.NewClient(k, t)

	bp := c.String("board")
	b, err := client.GetBoard(bp, trello.Defaults())
	if err != nil {
		return err
	}

	lists, _ := b.GetLists(trello.Defaults())
	cards, _ := lists[0].GetCards(trello.Defaults())

	pp.Print(lists[0])
	pp.Print(cards[1])

	return nil
}
