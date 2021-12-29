package handler

import (
	"fiber-101/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	var product []models.Product

	database.DBConn.Find(&product)

	return c.JSON(fiber.Map{
		"data": product,
	})
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DBConn
	var product models.Product
	db.Find(&product, id)
	if product.ID == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("Not found product id %v", id),
		})
	}

	return c.JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	errors := models.ValidateStruct(*product)
	if errors != nil {
		return c.JSON(errors)
	}

	db := database.DBConn
	db.Create(&product)
	return c.SendStatus(201)
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var product models.Product

	db.First(&product, id)

	if product.ID == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("Not found product id %v", id),
		})
	}

	db.Delete(&product)
	return c.SendStatus(204)
}
