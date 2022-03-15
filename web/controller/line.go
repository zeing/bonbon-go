package controller

import (
	"bonbon-go/service"
	lineservice "bonbon-go/service/line"
	"github.com/gin-gonic/gin"
)

type LineBotController struct {
	LineBotService lineservice.LineBotService
}

type LineInterfaceImpl interface {
	HandlerEvent(ctx *gin.Context) (string, error)
}

func NewLineBotController(services *service.Services) LineInterfaceImpl {
	return &LineBotController{
		LineBotService: services.LineBotService,
	}
}

func (svc *LineBotController) HandlerEvent(ctx *gin.Context) (string, error) {
	msg, err := svc.LineBotService.Handler(ctx)
	return msg, err
}
