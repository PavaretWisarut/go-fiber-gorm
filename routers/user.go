package routers

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-app/controllers"
)

func SetupUserRouter(app fiber.Router){

	app.Get("/users", controllers.GetUsers)
	app.Get("/user/:id", controllers.GetUser)
	app.Post("/user", controllers.SaveUser)
	app.Delete("/user/:id", controllers.DeleteUser)
	app.Put("/user/:id", controllers.UpdateUser)
}