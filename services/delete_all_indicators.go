package services

import (
	"crpalert/repositories"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DeleteAllIndicators(bot *tgbotapi.BotAPI, chatId int64) {
	indicators := repositories.DeleteAllIndicators(chatId)
	text := "Deleted indicators:\n"
	for i := 0; i < len(indicators); i++ {
		text += indicators[i].Indicator + " for " + indicators[i].Coin + " coin" + "\n"
	}

	msg := tgbotapi.NewMessage(chatId, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
