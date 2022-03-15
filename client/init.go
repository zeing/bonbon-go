package client

import (
	linebotclient "bonbon-go/client/lineclient"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Clients struct {
	LineBotClient *linebot.Client
}

func Init() *Clients {
	return &Clients{
		LineBotClient: linebotclient.NewLineBotClient(),
	}
}
