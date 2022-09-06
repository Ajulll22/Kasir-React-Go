package cart

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := app.Group("/api/cart")
	routes.Get("/", h.GetCart)
	routes.Post("/", h.AddCart)
	routes.Put("/:id", h.EditCart)
	routes.Delete("/:id", h.DeleteCart)
}
