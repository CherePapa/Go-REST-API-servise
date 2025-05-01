package routes

import (
	"web-service/controller"
	"web-service/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	app.Use(middleware.IsAuthentication)
	app.Post("/api/post", controller.CreatePost)
}
