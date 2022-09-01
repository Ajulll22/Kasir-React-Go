package product

import (
	"go-server/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetProduct(c *fiber.Ctx) error {
	if id := c.Query("id_product"); id != "" {

		var product models.Product

		h.DB.Where("id_product = ?", id).First(&product)

		return c.Status(fiber.StatusOK).JSON(&product)
	}
	if id := c.Query("id_category"); id != "" {

		var products []models.Product

		h.DB.Where("id_category = ?", id).First(&products)

		return c.Status(fiber.StatusOK).JSON(&products)
	}

	var products []models.Product

	result := h.DB.Find(&products)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&products)
}
