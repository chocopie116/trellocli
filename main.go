package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

	b := c.String("board")
	board, err := client.GetBoard(b, trello.Defaults())
	if err != nil {
		logger.Fatalln(errors.Wrap(err, fmt.Sprintf("failed Get Board by ID. Value: %#v", b)))
	}

	lists, err := board.GetLists(trello.Defaults())
	if err != nil {
		logger.Fatalln(errors.Wrap(err, fmt.Sprintf("failed Get List on Board. Value: %#v", board)))
	}

	fmt.Println(time.Now().Format("2006/01/02"))
	fmt.Println("```")
	for _, list := range lists {
		if list.Name == "TODO" {
			continue
		}

		cards, err := list.GetCards(trello.Defaults())
		if err != nil {
			logger.Fatalln(errors.Wrap(err, fmt.Sprintf("failed Get Cards on List. Value: %#v", list)))
		}

		fmt.Printf("# %s \n", list.Name)
		for _, card := range cards {
			fmt.Printf("- %s\n", card.Name)
		}
		fmt.Print("\n")
	}
	fmt.Println("```")

	return nil
}
