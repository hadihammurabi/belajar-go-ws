package main

import (
	"github.com/hadihammurabi/belajar-go-ws/config"
	"github.com/hadihammurabi/belajar-go-ws/internal/app/delivery/http"
	"github.com/hadihammurabi/belajar-go-ws/internal/app/delivery/ws"
	"github.com/hadihammurabi/belajar-go-ws/internal/app/repository"
	"github.com/hadihammurabi/belajar-go-ws/internal/app/service"
	"github.com/sarulabs/di"
)

// NewIOC func
func NewIOC(conf *config.Config) di.Container {
	builder, _ := di.NewBuilder()

	builder.Add(di.Def{
		Name: "config",
		Build: func(ctn di.Container) (interface{}, error) {
			return conf, nil
		},
	})

	builder.Add(di.Def{
		Name: "repository",
		Build: func(ctn di.Container) (interface{}, error) {
			return repository.NewRepository(builder.Build()), nil
		},
	})

	builder.Add(di.Def{
		Name: "service",
		Build: func(ctn di.Container) (interface{}, error) {
			return service.NewService(builder.Build()), nil
		},
	})

	deliveryHttp := http.Init(builder.Build())
	builder.Add(di.Def{
		Name: "delivery/http",
		Build: func(ctn di.Container) (interface{}, error) {
			return deliveryHttp, nil
		},
	})

	deliveryWs := ws.Init(builder.Build())
	builder.Add(di.Def{
		Name: "delivery/ws",
		Build: func(ctn di.Container) (interface{}, error) {
			return deliveryWs, nil
		},
	})

	return builder.Build()
}
