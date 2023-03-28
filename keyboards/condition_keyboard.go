package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ConditionKeyboard() tgbotapi.InlineKeyboardMarkup {
	var selectCoinKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("price is greater than", "condition=priceGreater"),
			tgbotapi.NewInlineKeyboardButtonData("price is less than", "condition=priceLess"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("price increases by %", "condition=increases%"),
			tgbotapi.NewInlineKeyboardButtonData("price decreases by %", "condition=decreases%"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("price spikes", "condition=spikes"),
			tgbotapi.NewInlineKeyboardButtonData("price drops", "condition=drops"),
		),
	)

	return selectCoinKeyboard
}
