package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SelectCoinKeyboard() tgbotapi.InlineKeyboardMarkup {
	var selectCoinKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("BTC", "coin=BTC"),
			tgbotapi.NewInlineKeyboardButtonData("ETH", "coin=ETH"),
			tgbotapi.NewInlineKeyboardButtonData("USDT", "coin=USDT"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("BNB", "coin=BNB"),
			tgbotapi.NewInlineKeyboardButtonData("USDC", "coin=USDC"),
			tgbotapi.NewInlineKeyboardButtonData("XRP", "coin=XRP"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("BUSD", "coin=BUSD"),
			tgbotapi.NewInlineKeyboardButtonData("ADA", "coin=ADA"),
			tgbotapi.NewInlineKeyboardButtonData("DOGE", "coin=DOGE"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("LINK", "coin=LINK"),
			tgbotapi.NewInlineKeyboardButtonData("ATOM", "coin=ATOM"),
			tgbotapi.NewInlineKeyboardButtonData("SHIB", "coin=SHIB"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Another symbol", "coin=type"),
		),
	)

	return selectCoinKeyboard
}
