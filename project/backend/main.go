package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
 
    "github.com/fmilheir/final_year_project/backend/database"
    "github.com/fmilheir/final_year_project/backend/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allows all origins
		AllowMethods: "GET,POST,PUT,PATCH,DELETE", // Specifies allowed methods
		AllowHeaders: "Origin, Content-Type, Accept", // Specifies allowed headers
		AllowCredentials: true, // Allows credentials
	}))

	routes.Setup(app)

	app.Listen(":8000")
}