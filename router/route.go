package router

import (
	"fiber-101/products"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/products", products.GetProducts)
	app.Get("/products/:id", products.GetProduct)
	app.Post("/products", products.CreateProduct)
	app.Delete("/products/:id", products.DeleteProduct)
}
