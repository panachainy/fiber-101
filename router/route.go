package router

import (
	"fiber-101/database"
	"fiber-101/products"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	p := products.NewProductRepo(database.DBConn)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// TODO: refactor other handlers to use the same pattern
	app.Get("/products", products.GetGetsFunc(p))
	app.Get("/products/:id", products.Get)
	app.Post("/products", products.Create)
	app.Delete("/products/:id", products.Delete)
}
