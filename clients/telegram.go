package clients

import (
	"crpalert/utils"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Init() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(utils.GetEnvVar("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	return bot
}
