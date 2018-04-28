package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adlio/trello"
	"github.com/pkg/errors"
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
	app.Run(os.Args)
}

func list(c *cli.Context) error {
	k := c.String("app_key")
	t := c.String("token")
	if k == "" || t == "" {
		logger.Fatalln("app_key and token is requred.")
	}

	client := trello.NewClient(k, t)

	bp := c.String("board")
	b, err := client.GetBoard(bp, trello.Defaults())
	if err != nil {
		logger.Fatalln(errors.Wrap(err, fmt.Sprintf("failed Get Board by ID. Value: %#v", bp)))
	}

	lists, err := b.GetLists(trello.Defaults())
	if err != nil {
		logger.Fatalln(errors.Wrap(err, fmt.Sprintf("failed Get List on Board. Value: %#v", b)))
	}

	for _, list := range lists {
		cards, err := list.GetCards(trello.Defaults())
		if err != nil {
			logger.Fatalln(errors.Wrap(err, fmt.Sprintf("failed Get Cards on List. Value: %#v", list)))
		}

		fmt.Printf("# %s \n", list.Name)
		for _, card := range cards {
			fmt.Printf("- %s\n", card.Name)
		}
		fmt.Print("\n\n")
	}

	return nil
}
