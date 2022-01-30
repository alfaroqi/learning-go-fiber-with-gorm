package route

import (
	"github.com/alfaroqi/learning-go-fiber-with-gorm/handler"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/user", handler.UserhandlerRead)
}
