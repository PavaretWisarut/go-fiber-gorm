package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-app/config"
	// "go-fiber-app/controllers"
	"go-fiber-app/routers"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.InitialMigration()
	app := fiber.New()

	corsConfig := cors.Config{
		AllowOrigins:     "*", // Change this to your desired list of allowed origins
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin,Content-Type,Accept",
		AllowCredentials: true,
	}
	app.Get("/", status)

	api := app.Group("/api")

	routers.SetupUserRouter(api)

	app.Use(cors.New(corsConfig))
	app.Listen(":3000")
}

func status(c *fiber.Ctx) error {
	return c.SendString("Server is running ! send your request ?")
}
