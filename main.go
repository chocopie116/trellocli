package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chocopie116/trellocli/util"

	"github.com/adlio/trello"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var logger = log.New(os.Stdout, "", log.Ldate+log.Ltime+log.Lshortfile)
var commands = []cli.Command{
	cli.Command{
		Name:   "list",
		Action: list,
	},
	cli.Command{
		Name: "add",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "card_name, c",
				Usage: "card name",
			},
		},
		Action: add,
	},
}

func main() {
	app := cli.NewApp()
	app.Commands = commands
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "config.toml",
			Usage: "target config toml path",
		},
	}
	app.Run(os.Args)
}

func initClient(c util.Config) *trello.Client {
	if c.ApiConfig.AppKey == "" || c.ApiConfig.Token == "" {
		logger.Fatalln("app_key and token is requred parameter.")
	}

	client := trello.NewClient(c.ApiConfig.AppKey, c.ApiConfig.Token)

	return client
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func list(c *cli.Context) error {
	path := c.GlobalString("config")
	config, err := util.ReadConfig(path)
	if err != nil {
		logger.Fatalln(errors.Wrap(err, fmt.Sprintf("cannot read config. Value: %s", path)))
	}

	client := initClient(config)
	board, err := client.GetBoard(config.TargetConfig.BoardId, trello.Defaults())
	if err != nil {
		logger.Fatalln(errors.Wrap(err, fmt.Sprintf("failed Get Board by ID. Value: %#v", config.TargetConfig.BoardId)))
	}

	lists, err := board.GetLists(trello.Defaults())
	if err != nil {
		logger.Fatalln(errors.Wrap(err, fmt.Sprintf("failed Get List on Board. Value: %#v", board)))
	}

	fmt.Println(time.Now().Format("2006/01/02"))
	fmt.Println("```")
	for _, list := range lists {
		if !contains(config.TargetConfig.ShowListNames, list.Name) {
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

func add(c *cli.Context) error {
	path := c.GlobalString("config")
	config, err := util.ReadConfig(path)
	if err != nil {
		logger.Fatalln(errors.Wrap(err, fmt.Sprintf("cannot read config. Value: %s", path)))
	}

	client := initClient(config)
	cn := c.String("card_name")
	if cn == "" {
		logger.Fatalln("card_name is required parameter")
	}

	card := trello.Card{
		Name:   cn,
		IDList: config.TargetConfig.AddListId,
	}

	err = client.CreateCard(&card, trello.Defaults())
	if err != nil {
		logger.Fatalln(errors.Wrap(err, "failed create Cards on List"))
	}

	return nil
}
