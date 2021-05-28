package router

import (
	"fiber-101/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/version", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	app.Get("/products", handler.GetAllProducts)
}
