package web

import (
	"bonbon-go/service"
	"bonbon-go/web/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func wrapper(f func(c *gin.Context) (string, error)) gin.HandlerFunc {

	return func(c *gin.Context) {
		_, err := f(c)
		if err != nil {
			c.JSON(503, gin.H{"status": err})
			return
		}
		c.JSON(200, gin.H{"status": "OK"})
	}
}

func Init(services *service.Services) *gin.Engine {
	srv := gin.Default()

	lineCtrl := controller.NewLineBotController(services)

	auth := srv.Group("/line")
	{
		auth.POST("/webhook", wrapper(lineCtrl.HandlerEvent))
	}

	srv.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "ohai")
	})

	return srv
}
