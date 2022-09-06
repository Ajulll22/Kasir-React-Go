package cart

import (
	"go-server/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) DeleteCart(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err == nil {
		var cart models.Cart
		err := h.DB.First(&cart, id).Error
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}

		h.DB.Delete(&cart)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Deleted",
		})
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Error Parameter",
	})
}
