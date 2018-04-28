package main

import (
	"fmt"
	"os"

	"github.com/adlio/trello"
	"github.com/urfave/cli"
)

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

	lists, err := b.GetLists(trello.Defaults())
	if err != nil {
		return err
	}

	for _, list := range lists {
		cards, err := list.GetCards(trello.Defaults())
		if err != nil {
			return err
		}

		fmt.Printf("%s \n", list.Name)
		for _, card := range cards {
			fmt.Printf("- %s\n", card.Name)
		}
		fmt.Print("\n\n")
	}

	return nil
}
