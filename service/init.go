package service

import (
	"bonbon-go/client"
	lineservice "bonbon-go/service/line"
)

type Services struct {
	LineBotService lineservice.LineBotService
}

func Init(cli *client.Clients) *Services {
	return &Services{
		LineBotService: lineservice.Init(cli.LineBotClient),
	}
}
