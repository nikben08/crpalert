package services

import (
	"crpalert/repositories"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ShowAllActiveIndicators(bot *tgbotapi.BotAPI, chatId int64) {
	rows := repositories.GetAllIndicators(chatId)
	text := "Active indicators:\n"
	for i := 0; i < len(rows); i++ {
		text += "coin: " + rows[i]["coin"] + " indicator: " + rows[i]["indicator"] + " interval: " + rows[i]["interval"] + "\n"
	}
	msg := tgbotapi.NewMessage(chatId, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
