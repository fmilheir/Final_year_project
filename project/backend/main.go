package main

import (
	//"log"
	//"os"
	//"fmt"

	//"github.com/go-openapi/loads"
	//flags "github.com/jessevdk/go-flags"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/fmilheir/final_year_project/backend/database"
	"github.com/fmilheir/final_year_project/backend/routes"
	//"github.com/fmilheir/final_year_project/backend/restapi"
	//"github.com/fmilheir/final_year_project/backend/restapi/operations"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		ExposeHeaders:    "set-cookie",
		
	}))

	routes.Setup(app)

	app.Listen(":8080")
}
