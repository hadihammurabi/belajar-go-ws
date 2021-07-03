package http

import (
	"fmt"
	"log"

	"github.com/hadihammurabi/belajar-go-ws/config"
	"github.com/hadihammurabi/belajar-go-ws/internal/app/delivery/http/middleware"
	"github.com/hadihammurabi/belajar-go-ws/internal/app/service"
	"github.com/sarulabs/di"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Delivery struct
type Delivery struct {
	HTTP        *fiber.App
	Middlewares func(int) fiber.Handler
	Service     *service.Service
	Validator   *config.Validator
	Config      *config.Config
}

// Init func
func Init(ioc di.Container) *Delivery {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[\"${time}\", \"${method}\", \"${path}\", \"${status}\", \"${ip}\", \"${latency}\"]\n",
	}))
	app.Use(recover.New())
	app.Use(cors.New())

	middleware.Middlewares = map[int]fiber.Handler{}
	middleware.Middlewares[middleware.AUTH] = middleware.NewAuthMiddleware(config.ConfigureJWT())

	service := ioc.Get("service").(*service.Service)
	conf := ioc.Get("config").(*config.Config)

	delivery := &Delivery{
		HTTP:        app,
		Middlewares: middleware.Use,
		Service:     service,
		Config:      conf,
		Validator:   config.NewValidator(),
	}
	delivery.ConfigureRoute()
	return delivery
}

func (d *Delivery) Run() {
	connURL := fmt.Sprintf(":%s", d.Config.APP.Port)
	if err := d.HTTP.Listen(connURL); err != nil {
		log.Printf("Server is not running! Reason: %v", err)
	}
}
