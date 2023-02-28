package handlers

import (
	"crpalert/repositories"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Messages(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatIdStr := strconv.FormatInt(int64(update.Message.Chat.ID), 10)
	status := repositories.WhetherChatExist(chatIdStr)
	if !status {
		repositories.CreateChat(chatIdStr)
	}

	if len(update.Message.Text) == 0 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please, enter correct data")
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}

}
