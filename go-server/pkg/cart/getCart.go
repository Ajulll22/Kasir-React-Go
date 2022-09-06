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

	err := h.DB.Joins("JOIN products on products.id_product = carts.id_product").
		Joins("JOIN categories on categories.id_category = products.id_category").
		Preload("Product.Category").
		Find(&carts).Error
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&carts)
}
