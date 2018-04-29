package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/BurntSushi/toml"
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
		},
		Action: list,
	},
	cli.Command{
		Name: "add",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "board",
				Usage: "Target board",
			},
			cli.StringFlag{
				Name:  "card_name",
				Usage: "card name",
			},
		},
		Action: add,
	},
}

type Config struct {
	ApiConfig    ApiConfig    `toml:"api"`
	TargetConfig TargetConfig `toml:"target"`
}

type ApiConfig struct {
	AppKey string `toml:"app_key"`
	Token  string `toml:"token"`
}
type TargetConfig struct {
	BoardId   string `toml:"board_id"`
	AddListId string `toml:"add_list_id"`
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

func initClient(path string) *trello.Client {
	var c Config
	_, err := toml.DecodeFile(path, &c)
	if err != nil {
		logger.Fatalln(errors.Wrap(err, fmt.Sprintf("cannot read config. Value: %s", path)))
	}

	if c.ApiConfig.AppKey == "" || c.ApiConfig.Token == "" {
		logger.Fatalln("app_key and token is requred parameter.")
	}

	client := trello.NewClient(c.ApiConfig.AppKey, c.ApiConfig.Token)

	return client
}

func list(c *cli.Context) error {
	configPath := c.GlobalString("config")
	client := initClient(configPath)

	b := c.String("board")
	if b == "" {
		logger.Fatalln("board is required parameter")
	}
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

func add(c *cli.Context) error {
	configPath := c.GlobalString("config")
	client := initClient(configPath)
	cn := c.String("card_name")
	if cn == "" {
		logger.Fatalln("card_name is required parameter")
	}

	card := trello.Card{
		Name:   cn,
		IDList: "59aabd6963041edd45105f05", //TODO ハードコーディングやめる
	}

	err := client.CreateCard(&card, trello.Defaults())
	if err != nil {
		logger.Fatalln(errors.Wrap(err, "failed create Cards on List"))
	}

	return nil
}
