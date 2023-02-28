package handlers

import (
	"crpalert/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Commands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		services.Start(bot, update)
	case "set_indicator":
		services.SetCoin(bot, update)
	case "show_all_indicators":
		services.ShowAllActiveIndicators(bot, update.Message.Chat.ID)
	case "delete_all_indicators":
		services.DeleteAllIndicators(bot, update.Message.Chat.ID)
	case "delete_indicator":
		services.DeleteIndicator(bot, update.Message.Chat.ID)
	}
}
