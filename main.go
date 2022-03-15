package main

import (
	"bonbon-go/client"
	"bonbon-go/config"
	"bonbon-go/service"
	"bonbon-go/web"
)

func main() {
	config.Init("")

	clients := client.Init()
	services := service.Init(clients)
	srv := web.Init(services)

	srv.Run(":" + config.App.Port)
}
