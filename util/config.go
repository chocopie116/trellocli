package util

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	ApiConfig    ApiConfig    `toml:"api"`
	TargetConfig TargetConfig `toml:"target"`
}

type ApiConfig struct {
	AppKey string `toml:"app_key"`
	Token  string `toml:"token"`
}
type TargetConfig struct {
	BoardId       string   `toml:"board_id"`
	AddListId     string   `toml:"add_list_id"`
	ShowListNames []string `toml:"show_list_names"`
	AddListName   string   `toml:"add_list_name"`
}

func ReadConfig(path string) (Config, error) {
	var c Config
	_, err := toml.DecodeFile(path, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}
