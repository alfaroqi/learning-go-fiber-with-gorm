package handler

import (
	"github.com/alfaroqi/learning-go-fiber-with-gorm/database"
	"github.com/alfaroqi/learning-go-fiber-with-gorm/model/entity"
	"github.com/gofiber/fiber/v2"
)

func UserhandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User

	result := database.DB.Debug().Find(&users)
	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})

	}
	return ctx.JSON(users)
}
