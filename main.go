package main

import (
	"fmt"

	"github.com/hadihammurabi/belajar-go-ws/config"
	deliveryHttp "github.com/hadihammurabi/belajar-go-ws/internal/app/delivery/http"
	"github.com/hadihammurabi/belajar-go-ws/internal/app/delivery/ws"

	_ "github.com/hadihammurabi/belajar-go-ws/docs"

	"github.com/joho/godotenv"
)

// @title Belajar Go REST API
// @version 0.0.1
// @description Ini adalah projek untuk latihan REST API dengan Go
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	_ = godotenv.Load()

	conf, err := config.New()
	if err != nil {
		panic(err)
	}

	ioc := NewIOC(conf)
	httpApp := deliveryHttp.Init(ioc)
	wsApp := ioc.Get("delivery/ws").(*ws.Delivery)

	forever := make(chan bool)
	go httpApp.Run()
	go wsApp.HTTP.Listen(fmt.Sprintf(":%s", conf.APP.WsPort))
	<-forever
}
