package web

import (
	"bonbon-go/service"
	"bonbon-go/web/controller"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func wrapperLineWebhook(f func(c *gin.Context) (string, error)) gin.HandlerFunc {

	return func(c *gin.Context) {
		_, err := f(c)
		if err != nil {
			log.Ctx(c).Error().Err(err).Msg("Wrap error")
		}
		c.JSON(200, gin.H{"status": "OK"})
	}
}

func Init(services *service.Services) *gin.Engine {
	srv := gin.Default()

	lineCtrl := controller.NewLineBotController(services)
	telegramCtrl := controller.NewTelegramController(services)

	auth := srv.Group("/line")
	{
		auth.POST("/webhook", wrapperLineWebhook(lineCtrl.HandlerEvent))
	}

	srv.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "ohai")
	})
	go func() {
		telegramCtrl.HandlerEvent()
	}()

	return srv
}
