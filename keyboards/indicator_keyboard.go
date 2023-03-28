package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func IndicatorKeyboard() tgbotapi.InlineKeyboardMarkup {
	var indicatorKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("MA", "indicator=ma"),
			tgbotapi.NewInlineKeyboardButtonData("EMA", "indicator=ema"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("MACD", "indicator=macd"),
			tgbotapi.NewInlineKeyboardButtonData("RSI", "indicator=rsi"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("PPO", "indicator=ppo"),
			tgbotapi.NewInlineKeyboardButtonData("Parabolic SAR", "indicator=sar"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("ADX", "indicator=adx"),
			tgbotapi.NewInlineKeyboardButtonData("STOCH", "indicator=stoch"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Standard deviation", "indicator=stddev"),
			tgbotapi.NewInlineKeyboardButtonData("Bollinger bands", "indicator=bbands"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Fibonacci retracement", "indicator=fibonacciretracement"),
			tgbotapi.NewInlineKeyboardButtonData("Williams’ %R", "indicator=willr"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("CCI", "indicator=cci"),
			tgbotapi.NewInlineKeyboardButtonData("Ichimoku cloud", "indicator=ichimoku"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("OBV", "indicator=obv"),
			tgbotapi.NewInlineKeyboardButtonData("Chaikin A/D Line", "indicator=ad"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Aroon", "indicator=aroon"),
			tgbotapi.NewInlineKeyboardButtonData("Aroon Oscillator", "indicator=aroonosc"),
		),
	)

	return indicatorKeyboard
}
