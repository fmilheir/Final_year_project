package routes

import (
	"github.com/fmilheir/final_year_project/backend/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register_admin)
	app.Post("/api/register_user", controller.Register_user)
	app.Post("/api/login_admin", controller.Login_admin)
	app.Post("/api/login_user", controller.Login_user)
	app.Get("/api/user", controller.User)
	app.Post("/api/logout", controller.Logout)
	//app.Get("/api/refresh", controller.Refresh)
}
