package bot

import (
	"fmt"
	"github-telegram-bot/pkg/config"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api         *tgbotapi.BotAPI
	githubToken string
}

func Start(cfg config.Config) {
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Panic(err)
	}

	b := &Bot{
		api:         bot,
		githubToken: cfg.GithubToken,
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			b.handleUpdate(update)
		}
	}
}
