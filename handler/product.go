package handler

import (
	"fiber-101/database"
	"fiber-101/model"

	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(c *fiber.Ctx) error {
	var product []model.Product

	database.DBConn.Find(&product)

	return c.JSON(fiber.Map{
		"message": "I'm a GET request!",
		"data":    product,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	product := new(model.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	errors := model.ValidateStruct(*product)
	if errors != nil {
		return c.JSON(errors)
	}

	db := database.DBConn
	db.Create(&product)
	return c.SendStatus(201)
}
