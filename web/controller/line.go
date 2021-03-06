package controller

import (
	"bonbon-go/service"
	lineservice "bonbon-go/service/line"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type LineInterfaceImpl struct {
	LineBotService lineservice.LineBotService
}

type LineBotController interface {
	HandlerEvent(ctx *gin.Context) (string, error)
}

func NewLineBotController(services *service.Services) LineBotController {
	return &LineInterfaceImpl{
		LineBotService: services.LineBotService,
	}
}

func (svc *LineInterfaceImpl) HandlerEvent(ctx *gin.Context) (string, error) {
	event, tweet, err := svc.LineBotService.Handler(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error => Line Event Error")
		return "", err
	}
	if tweet != nil {
		replyMessage := fmt.Sprintf(
			"Tweeted !! | See at https://twitter.com/bon2_official/status/%s", tweet.IDStr)
		err = svc.LineBotService.ReplyToUser(event, replyMessage)
		if err != nil {
			log.Error().Err(err).Msg("Error => Twitter Tweet Error")
			return "", err
		}
	}

	return "", nil
}
