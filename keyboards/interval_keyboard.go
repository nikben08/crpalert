package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func IntervalKeyboard() tgbotapi.InlineKeyboardMarkup {
	var indicatorKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1m", "interval=1m"),
			tgbotapi.NewInlineKeyboardButtonData("5m", "interval=5m"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("15m", "interval=15m"),
			tgbotapi.NewInlineKeyboardButtonData("30m", "interval=30m"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1h", "interval=1h"),
			tgbotapi.NewInlineKeyboardButtonData("2h", "interval=2h"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("4h", "interval=4h"),
			tgbotapi.NewInlineKeyboardButtonData("12h", "interval=12h"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1d", "interval=1d"),
			tgbotapi.NewInlineKeyboardButtonData("1w", "interval=1w"),
		),
	)

	return indicatorKeyboard
}
