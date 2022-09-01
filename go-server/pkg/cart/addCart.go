package cart

import (
	"go-server/config"
	"go-server/models"

	"github.com/gofiber/fiber/v2"
)

type CartRequest struct {
	Id_product string `json:"id_product" validate:"required,number"`
}

func (h handler) AddCart(c *fiber.Ctx) error {
	data := new(CartRequest)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}

	errors := config.Validate(*data)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	var product models.Product
	h.DB.Where("id_product = ?", data.Id_product).First(&product)

	var cart models.Cart
	h.DB.Where("id_product = ? AND status_cart = ?", data.Id_product, 1).First(&cart)
	if cart.Id_cart != 0 {
		cart.Jumlah = cart.Jumlah + 1
		cart.Total_harga = cart.Jumlah * product.Harga_product

		h.DB.Save(&cart)
		return c.Status(fiber.StatusOK).JSON(&cart)
	}

	cart.Jumlah = 1
	cart.Total_harga = product.Harga_product
	cart.Id_product = product.Id_product
	h.DB.Create(&cart)

	return c.Status(fiber.StatusCreated).JSON(&cart)
}
