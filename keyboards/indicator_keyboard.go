package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func IndicatorKeyboard() tgbotapi.InlineKeyboardMarkup {
	var indicatorKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("MA", "indicator=MA"),
			tgbotapi.NewInlineKeyboardButtonData("EMA", "indicator=EMA"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("MACD", "indicator=MACD"),
			tgbotapi.NewInlineKeyboardButtonData("RSI", "indicator=RSI"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("PPO", "indicator=PPO"),
			tgbotapi.NewInlineKeyboardButtonData("Parabolic SAR", "indicator=SAR"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("ADX", "indicator=ADX"),
			tgbotapi.NewInlineKeyboardButtonData("Stochastic oscillator", "indicator=stochastic_oscillator"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Standard deviation", "indicator=standard_deviation"),
			tgbotapi.NewInlineKeyboardButtonData("Bollinger bands", "indicator=bollinger_bands"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Fibonacci retracement", "indicator=fibonacci_retracement"),
			tgbotapi.NewInlineKeyboardButtonData("Williams %R", "indicator=williams_r"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("CCI", "indicator=CCI"),
			tgbotapi.NewInlineKeyboardButtonData("Ichimoku cloud", "indicator=ichimoku_cloud"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("OBV", "indicator=OBV"),
			tgbotapi.NewInlineKeyboardButtonData("A/D line", "indicator=A/D_line"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Aroon indicator", "indicator=aroon"),
		),
	)

	return indicatorKeyboard
}
