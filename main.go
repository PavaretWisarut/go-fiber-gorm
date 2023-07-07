package main

import (
	"go-fiber-app/config"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	// "go-fiber-app/controllers"
	"go-fiber-app/routers"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
	routers.SetupProviderRouter(api)

	app.Use(cors.New(corsConfig))
	app.Listen(":3000")
}

func status(c *fiber.Ctx) error {
	return c.SendString("Server is running ! send your request ?")
}
