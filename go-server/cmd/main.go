package main

import (
	"log"

	"go-server/config"
	"go-server/db"
	"go-server/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed to load config ", err)
	}

	app := fiber.New()

	db := db.Conect(c.DBUrl)
	db.AutoMigrate(&models.Category{}, &models.Product{}, &models.Cart{}, &models.Transaction{})

	app.Listen(c.Port)
}
