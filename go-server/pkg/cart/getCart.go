package cart

import (
	"go-server/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetCart(c *fiber.Ctx) error {
	if id := c.Query("id"); id != "" {

		var cart models.Cart

		h.DB.Where("id_cart = ?", id).First(&cart)

		return c.Status(fiber.StatusOK).JSON(&cart)
	}
	var carts []models.Cart

	result := h.DB.Find(&carts)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&carts)
}
