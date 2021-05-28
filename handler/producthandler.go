package handler

import (
	"fiber-101/database"
	"fiber-101/model"

	"github.com/gofiber/fiber/v2"
)

// GetAllProducts from db
func GetAllProducts(c *fiber.Ctx) error {
	var product []model.Product
	db := database.DBConn
	db.Find(&product)

	return c.JSON(fiber.Map{
		"message": "I'm a GET request!",
		"data":    product,
	})
}

// GetSingleProduct from db
func GetSingleProduct(c *fiber.Ctx) {
}

// database.DBConn.Create(&model.Product{Code: "D42", Price: 100})
