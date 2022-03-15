package client

import (
	linebotclient "bonbon-go/client/lineclient"
	"bonbon-go/client/twitterclient"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Clients struct {
	LineBotClient *linebot.Client
	TwitterClient twitterclient.TwitterClient
}

func Init() *Clients {
	return &Clients{
		LineBotClient: linebotclient.NewLineBotClient(),
		TwitterClient: twitterclient.NewTwitterClient(),
	}
}
