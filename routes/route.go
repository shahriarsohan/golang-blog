package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shahriarsohan/go-blog-practise/controllers"
	"github.com/shahriarsohan/go-blog-practise/middleware"
)

func SetUp(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middleware.IsAuthenticated) // *** any routes below this line, will need to be authenticated
	app.Post("/api/blog/create", controllers.CreatePost)
	app.Get("/api/blog/list", controllers.AllPost)
	app.Get("/api/blog/d/:id", controllers.PostDetails)
	app.Put("/api/blog/d/:id", controllers.UpdatePost)
	app.Get("/api/blog/user/list", controllers.GetUserPost)
	app.Delete("/api/blog/delete/:id", controllers.DeletePost)
	app.Delete("/api/upload-image", controllers.Upload)

}
