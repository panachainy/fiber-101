package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Todo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{ID: 1, Name: "abc", Completed: false},
	{ID: 2, Name: "def", Completed: true},
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/version", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	app.Listen(":5000")
}
