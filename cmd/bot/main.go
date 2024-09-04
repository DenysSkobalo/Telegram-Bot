package main

import (
	"github-telegram-bot/pkg/bot"
	"github-telegram-bot/pkg/config"
)

func main() {
	cfg := config.LoadConfig()
	bot.Start(cfg)
}