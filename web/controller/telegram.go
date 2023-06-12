package controller

import (
	"bonbon-go/service"
	telegramservice "bonbon-go/service/telegram"
)

type TelegramInterfaceImpl struct {
	TelegramService telegramservice.TelegramService
}

type TelegramController interface {
	HandlerEvent()
}

func NewTelegramController(services *service.Services) TelegramController {
	return &TelegramInterfaceImpl{
		TelegramService: services.TelegramService,
	}
}

func (tic *TelegramInterfaceImpl) HandlerEvent() {
	tic.TelegramService.Handler()
}
