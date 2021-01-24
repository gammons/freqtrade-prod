package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Exchange struct {
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

type Telegram struct {
	Token  string `json:"token"`
	ChatID string `json:"chat_id"`
}

type Config struct {
	Exchange *Exchange `json:"exchange"`
	Telegram *Telegram `json:"telegram"`
}

func main() {
	config := &Config{
		Exchange: &Exchange{
			Key:    os.Getenv("BINANCE_EXCHANGE_KEY"),
			Secret: os.Getenv("BINANCE_EXCHANGE_SECRET"),
		},
		Telegram: &Telegram{
			Token:  os.Getenv("TELEGRAM_TOKEN"),
			ChatID: os.Getenv("TELEGRAM_CHAT_ID"),
		},
	}

	marshalled, _ := json.Marshal(config)
	ioutil.WriteFile("config.secrets.json", []byte(marshalled), 0644)
}
