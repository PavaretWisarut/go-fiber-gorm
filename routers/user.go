package routers

import (
	"go-fiber-app/controllers"
	"go-fiber-app/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRouter(app fiber.Router) {

	app.Get("/users", middleware.Authmiddleware(), controllers.GetUsers)
	app.Get("/user/:id", middleware.Authmiddleware(), controllers.GetUser)
	app.Post("/user", middleware.Authmiddleware(), controllers.SaveUser)
	app.Delete("/user/:id", middleware.Authmiddleware(), controllers.DeleteUser)
	app.Put("/user/:id", middleware.Authmiddleware(), controllers.UpdateUser)
}
