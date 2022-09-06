package cart

import (
	"go-server/config"
	"go-server/models"

	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type EditRequest struct {
	Jumlah     json.Number `json:"jumlah" validate:"required,number"`
	Keterangan string      `json:"keterangan"`
}

func (h handler) EditCart(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err == nil {
		data := new(EditRequest)

		if err := c.BodyParser(data); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})

		}

		errors := config.Validate(*data)
		if errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors)

		}

		var cart models.Cart
		err := h.DB.Preload("Product").First(&cart, id).Error
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}

		jumlah, _ := data.Jumlah.Int64()
		cart.Jumlah = int(jumlah)
		cart.Total_harga = cart.Product.Harga_product * int(jumlah)
		cart.Keterangan = data.Keterangan

		h.DB.Save(&cart)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Updated",
		})
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Error Parameter",
	})
}
