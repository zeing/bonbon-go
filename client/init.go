package client

import (
	linebotclient "bonbon-go/client/lineclient"
	"bonbon-go/client/telegramclient"
	"bonbon-go/client/twitterclient"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Clients struct {
	LineBotClient  *linebot.Client
	TwitterClient  twitterclient.TwitterClient
	TelegramClient *tgbotapi.BotAPI
}

func Init() *Clients {
	return &Clients{
		LineBotClient:  linebotclient.NewLineBotClient(),
		TwitterClient:  twitterclient.NewTwitterClient(),
		TelegramClient: telegramclient.NewTelegramClient(),
	}
}
