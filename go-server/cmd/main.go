package main

import (
	"log"

	"go-server/config"
	"go-server/db"
	"go-server/pkg/cart"
	"go-server/pkg/category"
	"go-server/pkg/product"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed to load config ", err)
	}

	app := fiber.New()
	app.Use(cors.New())

	db := db.Conect(c.DBUrl)
	// db.AutoMigrate(&models.Category{}, &models.Product{}, &models.Cart{}, &models.Transaction{})

	product.RegisterRoutes(app, db)
	category.RegisterRoutes(app, db)
	cart.RegisterRoutes(app, db)

	app.Listen(c.Port)
}
