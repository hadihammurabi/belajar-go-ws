package ws

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type ChatClient struct{}
type ChatMessage struct {
	Connection *websocket.Conn
	Body       string
}

type ChatHub struct {
	Clients    map[*websocket.Conn]ChatClient
	Register   chan *websocket.Conn
	Unregister chan *websocket.Conn
	Broadcast  chan ChatMessage
}

func (delivery *Delivery) InitChat() *ChatHub {
	hub := NewChatHub()
	go hub.Run()

	delivery.HTTP.Get("/mantap", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": "mantap",
		})
	})

	delivery.HTTP.Get("/ws", websocket.New(func(c *websocket.Conn) {
		defer func() {
			hub.Unregister <- c
			c.Close()
		}()

		hub.Register <- c

		for {
			_, in, err := c.ReadMessage()
			if err != nil {
				log.Println("err:in:", err)
				return
			}
			log.Println("res:in:", string(in))

			out := ChatMessage{
				Connection: c,
				Body:       fmt.Sprintf("Got \"%s\"", in),
			}
			hub.Broadcast <- out
			log.Println("res:out:", out)
		}
	}))

	return hub
}

func NewChatHub() *ChatHub {
	return &ChatHub{
		Clients:    make(map[*websocket.Conn]ChatClient),
		Register:   make(chan *websocket.Conn),
		Unregister: make(chan *websocket.Conn),
		Broadcast:  make(chan ChatMessage),
	}
}

func (h *ChatHub) Run() {
	for {
		select {
		case connection := <-h.Register:
			h.Clients[connection] = ChatClient{}
		case connection := <-h.Unregister:
			delete(h.Clients, connection)
		case message := <-h.Broadcast:
			for connection := range h.Clients {
				if message.Connection != connection {
					if err := connection.WriteMessage(websocket.TextMessage, []byte(message.Body)); err != nil {
						log.Printf("message broadcast error: %s", err)
						connection.WriteMessage(websocket.CloseMessage, []byte(err.Error()))
						connection.Close()
						delete(h.Clients, connection)
					}
				}
			}
		}
	}
}
