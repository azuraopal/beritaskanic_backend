package routes

import (
	"github.com/ValSpp/ber1taskanic/controller"
	"github.com/ValSpp/ber1taskanic/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Use(middleware.IsAuthenticate)

	app.Post("/api/post", controller.CreatePost)
	app.Get("/api/allpost", controller.AllPost)

}
