package category

import (
	"go-server/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetCategory(c *fiber.Ctx) error {
	if id := c.Query("id"); id != "" {

		var category models.Category

		h.DB.Where("id_category = ?", id).First(&category)

		return c.Status(fiber.StatusOK).JSON(&category)
	}
	var categories []models.Category

	result := h.DB.Find(&categories)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&categories)
}
