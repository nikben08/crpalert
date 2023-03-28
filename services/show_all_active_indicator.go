package services

import (
	"crpalert/repositories"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ShowAllIndicators(bot *tgbotapi.BotAPI, chatId int64) {
	indicators := repositories.GetAllIndicators(chatId)
	if len(indicators) != 0 {
		text := "Active indicators:\n"
		for i := 0; i < len(indicators); i++ {
			text += "coin: " + indicators[i].Coin + " indicator: " + indicators[i].Indicator + " interval: " + indicators[i].Frame + "\n"
		}
		msg := tgbotapi.NewMessage(chatId, text)
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
}
