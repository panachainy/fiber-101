package handler

import (
	"fiber-101/database"
	"fiber-101/model"

	"github.com/gofiber/fiber/v2"
)

// GetAllProducts from db
func GetAllProducts(c *fiber.Ctx) error {
	var product []model.Product

	database.DBConn.Find(&product)

	return c.JSON(fiber.Map{
		"message": "I'm a GET request!",
		"data":    product,
	})
}

// GetSingleProduct from db
func CreateProduct(c *fiber.Ctx) error {
	product := new(model.Product)

	if err := c.BodyParser(product); err != nil {
		return err
	}

	db := database.DBConn
	db.Create(&product)
	return c.SendStatus(201)
}