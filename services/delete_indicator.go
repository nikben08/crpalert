package services

import (
	"crpalert/repositories"
	"crpalert/utils"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DeleteIndicator(bot *tgbotapi.BotAPI, chatId int64) {
	data := repositories.GetAllIndicators(chatId)
	var btns []tgbotapi.InlineKeyboardButton
	for i := 0; i < len(data); i++ {
		btn := tgbotapi.NewInlineKeyboardButtonData("coin "+data[i].Coin+" indicator "+data[i].Indicator+" interval "+data[i].Frame, "delete_indicator="+data[i].Id)
		btns = append(btns, btn)
	}

	var rows [][]tgbotapi.InlineKeyboardButton
	for i := 0; i < len(btns); i += 2 {
		if i < len(btns) && i+1 < len(btns) {
			row := tgbotapi.NewInlineKeyboardRow(btns[i], btns[i+1])
			rows = append(rows, row)
		} else if i < len(btns) {
			row := tgbotapi.NewInlineKeyboardRow(btns[i])
			rows = append(rows, row)
		}
	}
	fmt.Println(len(rows))
	var keyboard = tgbotapi.InlineKeyboardMarkup{InlineKeyboard: rows}
	//keyboard.InlineKeyboard = rows

	text := "Please, select indicator you want to delete"
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ReplyMarkup = keyboard
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func DeleteIndicatorCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatId := update.CallbackQuery.Message.Chat.ID
	_, indicatorId := utils.GetKeyValue(update.CallbackQuery.Data)
	repositories.DeleteIndicator(indicatorId)
	text := "Indicator successfully deleted"
	msg := tgbotapi.NewMessage(chatId, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
