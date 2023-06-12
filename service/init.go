package service

import (
	"bonbon-go/client"
	lineservice "bonbon-go/service/line"
	telegramservice "bonbon-go/service/telegram"
)

type Services struct {
	LineBotService  lineservice.LineBotService
	TelegramService telegramservice.TelegramService
}

func Init(cli *client.Clients) *Services {
	return &Services{
		LineBotService:  lineservice.Init(cli.LineBotClient, cli.TwitterClient),
		TelegramService: telegramservice.Init(cli.TelegramClient, cli.TwitterClient),
	}
}
