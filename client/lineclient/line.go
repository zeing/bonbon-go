package linebotclient

import (
	"bonbon-go/config"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/rs/zerolog/log"
)

func NewLineBotClient() *linebot.Client {
	client, err := linebot.New(
		config.App.Line.ChannelSecret,
		config.App.Line.ChannelToken,
	)
	if err != nil {
		log.Error().Err(err).Msg("Line Client Error")
		panic(err)
	}
	return client
}
