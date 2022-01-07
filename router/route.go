package router

import (
	"fiber-101/products"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/products", products.Gets)
	app.Get("/products/:id", products.Get)
	app.Post("/products", products.Create)
	app.Delete("/products/:id", products.Delete)
}
