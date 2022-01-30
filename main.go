package main

import (
	"log"

	"github.com/alfaroqi/learning-go-fiber-with-gorm/database"
	"github.com/alfaroqi/learning-go-fiber-with-gorm/route"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {
	// initial database
	database.DatabaseInit()
	app := fiber.New()

	// inittialize routes
	route.RouteInit(app)

	app.Listen(":3000")

}
