package build

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var (
	Time    string
	User    string
	Version = "development"
)

func printBuildDetail() {
	logrus.Infoln("[BUILD] Version: ", Version)
	logrus.Infoln("[BUILD] build.Time: ", Time)
	logrus.Infoln("[BUILD] build.User: ", User)
}

func SetupVersion(app *fiber.App) {
	printBuildDetail()

	app.Get("/version", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"version": Version,
		})
	})
}
