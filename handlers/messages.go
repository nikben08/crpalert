package handlers

import (
	"crpalert/repositories"
	"crpalert/services"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Messages(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	//chatIdStr := strconv.FormatInt(int64(update.Message.Chat.ID), 10)
	// Message data validation
	if len(update.Message.Text) == 0 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please, enter correct data")
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	} else {
		msgType := repositories.SpecifyMsgType(update.Message.Chat.ID, update.Message.Text)
		if msgType == "symbol" {
			validateSymbol := repositories.ValidateSymbol(update.Message.Text)
			if validateSymbol {
				fmt.Println("Correct")
				services.SetTypedCoinCallback(bot, update.Message.Text, update.Message.Chat.ID)
			} else {
				fmt.Println("wrong")
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Incorrect crypto symbol")
				if _, err := bot.Send(msg); err != nil {
					panic(err)
				}
			}
		}
	}

}
