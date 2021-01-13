package config

import (
	"encoding/json"
	"os"
)

type BotConfig struct {
	Token          string   `json:"token"`
	Prefix         string   `json:"prefix"`
	CommandChannel string   `json:"command_channel"`
	MusicChannel   string   `json:"music_channel"`
	Database       database `json:"database"`
}

type database struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Host     string `json:"host"`
}

func Parse() (c *BotConfig, err error) {
	f, err := os.Open("./config/config.json")
	if err != nil {
		return
	}
	c = new(BotConfig)
	err = json.NewDecoder(f).Decode(c)
	return
}
