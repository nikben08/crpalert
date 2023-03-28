package services

import (
	"crpalert/keyboards"
	"crpalert/repositories"
	"crpalert/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SetCoin(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	repositories.CreateEmptyIndicator(update.Message.Chat.ID)
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
	repositories.SetCoin(update.CallbackQuery.Message.Chat.ID, coin)
}

func TypeCoinCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Please, type symbol you want"
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func SetTypedCoinCallback(bot *tgbotapi.BotAPI, symbol string, chatId int64) {
	text := "Please, select indicator you want to track"
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ReplyMarkup = keyboards.IndicatorKeyboard()
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
	repositories.SetCoin(chatId, symbol)
}

func SetIndicatorCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	_, indicator := utils.GetKeyValue(update.CallbackQuery.Data)
	repositories.SetIndicator(update.CallbackQuery.Message.Chat.ID, indicator)
	text := "Please, select dataframe"
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboards.IntervalKeyboard()
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func SetFrameCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	_, frame := utils.GetKeyValue(update.CallbackQuery.Data)
	repositories.SetFrame(update.CallbackQuery.Message.Chat.ID, frame)
	text := "Please, select condition"
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboards.ConditionKeyboard()
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func SetConditionCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	_, condition := utils.GetKeyValue(update.CallbackQuery.Data)
	repositories.SetCondition(update.CallbackQuery.Message.Chat.ID, condition)
	text := "Indicator successfully set"
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
