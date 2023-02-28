package services

import (
	"crpalert/repositories"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DeleteAllIndicators(bot *tgbotapi.BotAPI, chatId int64) {
	rows := repositories.GetAllIndicators(chatId)
	text := "Deleted indicators:\n"
	for i := 0; i < len(rows); i++ {
		text += rows[i]["indicator"] + " for " + rows[i]["coin"] + " coin" + "\n"
	}

	repositories.DeleteAllIndicators(chatId)

	msg := tgbotapi.NewMessage(chatId, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
