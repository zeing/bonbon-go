package telegramclient

import (
	"bonbon-go/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

func NewTelegramClient() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(config.App.Telegram.Token)
	if err != nil {
		log.Panic().Err(err).Msg("Telegram Client error")
	}

	bot.Debug = false
	log.Info().Msgf("Authorized on account %s", bot.Self.UserName)

	return bot
}
