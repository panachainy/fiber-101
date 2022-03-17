package product

import (
	"fmt"

	"fiber-101/database"

	"github.com/gofiber/fiber/v2"
)

func GetGetsFunc(p ProductRepository) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		products := p.GetAll()

		return c.JSON(fiber.Map{
			"data": products,
		})
	}
}

func Get(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DBConn
	var product Product
	db.Find(&product, id)
	if product.ID == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("Not found product id %v", id),
		})
	}

	return c.JSON(product)
}

func Create(c *fiber.Ctx) error {
	product := new(Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	errors := ValidateStruct(*product)
	if errors != nil {
		return c.JSON(errors)
	}

	db := database.DBConn
	db.Create(&product)
	return c.SendStatus(201)
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var product Product

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
