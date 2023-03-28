package handlers

import (
	"crpalert/services"
	"crpalert/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Delete(c tgbotapi.Chattable, bot *tgbotapi.BotAPI) {
	bot.Send(c)
}

func Callbacks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	if _, err := bot.Request(callback); err != nil {
		panic(err)
	}

	cmd, value := utils.GetKeyValue(update.CallbackQuery.Data)

	switch {
	case cmd == "coin" && value != "type":
		services.SetCoinCallback(bot, update)
	case cmd == "coin" && value == "type":
		services.TypeCoinCallback(bot, update)
	case cmd == "indicator":
		services.SetIndicatorCallback(bot, update)
	case cmd == "interval":
		services.SetFrameCallback(bot, update)
	case cmd == "condition":
		services.SetConditionCallback(bot, update)
	case cmd == "delete_indicator":
		services.DeleteIndicatorCallback(bot, update)
	}

	msgId := update.CallbackQuery.Message.MessageID
	q := tgbotapi.NewDeleteMessage(update.CallbackQuery.Message.Chat.ID, msgId)
	Delete(q, bot)
}
