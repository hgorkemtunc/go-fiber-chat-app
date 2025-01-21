package main

import (
	"go-fiber-chat-app/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "./public")

	chatHandler := handlers.NewChatHandler()
	app.Get("/ws", websocket.New(chatHandler.HandleWebSocket))

	app.Listen(":3000")
}
