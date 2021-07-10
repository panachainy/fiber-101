package main

import (
	"fiber-101/build"
	"fiber-101/database"
	"fiber-101/router"
	"fiber-101/utils"
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
)

var (
	log  *logrus.Logger
	port = flag.String("port", ":5000", "Port to listen on")
)

func main() {
	app := SetupApp()

	log.Fatal(app.Listen(*port))
}

func SetupApp() *fiber.App {
	log = utils.InitLogrus()

	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's an fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			// Send custom error page
			err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(500).SendString("Internal Server Error")
			}

			// Return from handler
			return nil
		},
	})

	database.Init()

	app.Use(logger.New())
	app.Use(recover.New())

	build.SetupVersion(app)

	router.SetupRoutes(app)

	return app
}
