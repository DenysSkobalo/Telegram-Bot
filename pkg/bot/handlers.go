package bot

import "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *Bot) handleUpdate(update tgbotapi.Update) {
	switch update.Message.Command() {
	case "finduser":
			// b.handleFindUser(update)
	case "findrepo":
			// b.handleFindRepo(update)
	case "userrepos":
			// b.handleUserRepos(update)
	default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Unknown command")
			b.api.Send(msg)
	}
}