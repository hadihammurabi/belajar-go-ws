package http

import (
	"fmt"

	"github.com/hadihammurabi/belajar-go-ws/internal/app/delivery/http/middleware"
	"github.com/hadihammurabi/belajar-go-ws/internal/app/entity"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// NewAuthHandler func
func NewAuthHandler(delivery *Delivery) {
	router := delivery.HTTP.Group("/auth")
	router.Post("/login", delivery.Login)
	router.Get("/info", delivery.Middlewares(middleware.AUTH), delivery.Info)
}

// Login func
func (delivery Delivery) Login(c *fiber.Ctx) error {
	userInput := &entity.UserLoginDTO{}
	if err := c.BodyParser(userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	user := &entity.User{
		Email:    userInput.Email,
		Password: userInput.Password,
	}

	token, err := delivery.Service.Auth.Login(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"token": token,
	})
}

// Info func
func (delivery Delivery) Info(c *fiber.Ctx) error {
	fromLocals := c.Locals("user").(*jwt.Token)
	user, err := delivery.Service.JWT.GetUser(fromLocals.Raw)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	email := user.Email
	return c.SendString(fmt.Sprintf("welcome, %s", email))
}
