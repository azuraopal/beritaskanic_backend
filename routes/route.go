package routes

import (
	"github.com/ValSpp/ber1taskanic/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)

}
