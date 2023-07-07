package routers

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-app/controllers"
	
)

func SetupProviderRouter(app fiber.Router){
	
	app.Post("/register" , controllers.Register)
	app.Post("/login",controllers.Login)
}