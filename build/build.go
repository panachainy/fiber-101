package build

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

var (
	Time    string
	User    string
	Version = "development"
)

func printBuildDetail() {
	log.Println("Version:\t", Version)
	log.Println("build.Time:\t", Time)
	log.Println("build.User:\t", User)
}

func SetupVersion(app *fiber.App) {
	printBuildDetail()

	app.Get("/version", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"version": Version,
		})
	})
}
