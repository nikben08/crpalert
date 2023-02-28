package services

import (
	"crpalert/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Hi! This bot tracks changes in crypto coin indicators. " +
		"When changing, the bot will send you a notification." +
		"Please select the crypto coin you want to track."

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboards.CmdKeyboard()
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
