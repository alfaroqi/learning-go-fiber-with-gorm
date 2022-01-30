package main

import (
	"github.com/alfaroqi/learning-go-fiber-with-gorm/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// inittialize routes
	route.RouteInit(app)

	app.Listen(":3000")

}
