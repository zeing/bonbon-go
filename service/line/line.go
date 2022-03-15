package lineservice

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/rs/zerolog/log"
)

type defaultServices struct {
	lbc *linebot.Client
}

type LineBotService interface {
	Handler(ctx *gin.Context) (string, error)
}

func Init(LineBot *linebot.Client) LineBotService {
	return &defaultServices{
		lbc: LineBot,
	}
}

func (svc *defaultServices) Handler(ctx *gin.Context) (string, error) {
	events, err := svc.lbc.ParseRequest(ctx.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			return "", err
		} else {
			return "", err
		}
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if _, err = svc.lbc.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					log.Logger.Print(err)
				}
			case *linebot.StickerMessage:
				replyMessage := fmt.Sprintf(
					"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
				if _, err = svc.lbc.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
	return "", nil
}
