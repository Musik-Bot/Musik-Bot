package config

import (
	"encoding/json"
	"os"
)

type BotConfig struct {
	Token     string `json:"token"`
	Prefix    string `json:"prefix"`
	ChannelID string `json:"channel_id"`
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
