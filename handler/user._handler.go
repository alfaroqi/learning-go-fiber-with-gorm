package handler

import "github.com/gofiber/fiber/v2"

func UserhandlerRead(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "Hello World",
	})
	return nil
}
