package services

import (
	"crpalert/keyboards"
	"crpalert/repositories"
	"crpalert/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SetCoin(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Please, select coin"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboards.SelectCoinKeyboard()
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}

}

func SetCoinCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Please, select indicator you want to track"
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboards.IndicatorKeyboard()
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
	_, coin := utils.GetKeyValue(update.CallbackQuery.Data)
	repositories.CreateSetIndicatorCmd(update.CallbackQuery.Message.Chat.ID, coin)
}

func SetIndicatorCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	_, indicator := utils.GetKeyValue(update.CallbackQuery.Data)
	repositories.AddIndicatorForSetIndicatorCmd(update.CallbackQuery.Message.Chat.ID, indicator)
	text := "Please, select dataframe"
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboards.IntervalKeyboard()
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func SetFrameCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	_, frame := utils.GetKeyValue(update.CallbackQuery.Data)
	repositories.AddFrameForSetIndicatorCmd(update.CallbackQuery.Message.Chat.ID, frame)
	text := "Indicator successfully set"
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
