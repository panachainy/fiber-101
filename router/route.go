package router

import (
	"fiber-101/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/graphql"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/products", handler.GetProducts)

	app.Get("/products/:id", handler.GetProduct)

	app.Post("/products", handler.CreateProduct)

	app.Delete("/products/:id", handler.DeleteProduct)

	// queryType := graphql.NewObject(graphql.ObjectConfig{
	// 	// ...
	// })

	// Schema, _ := graphql.NewSchema(graphql.SchemaConfig{
	// 	// ...
	// })

	// h := gqlhandler.New(&gqlhandler.Config{
	// 	Schema: &Schema,
	// 	Pretty: true,
	// })

	// app.Handler("/graphql", h)
}
