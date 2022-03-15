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

	port := config.App.Port
	if port == "" {
		port = "8080"
	}
	srv.Run(":" + port)
}
