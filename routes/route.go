package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shahriarsohan/go-blog-practise/controllers"
)

func SetUp(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
}
