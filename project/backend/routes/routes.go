package routes

import (
	"github.com/fmilheir/final_year_project/backend/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register_admin", controller.Register_admin)
	app.Post("/api/getCompanyId", controller.GetCompanyID)
	app.Post("/api/register_user", controller.Register_user)
	app.Post("/api/login_admin", controller.Login_admin)
	app.Post("/api/login_user", controller.Login_user)
	app.Post("/api/chatbot", controller.MyHandler)
	app.Post("/api/update_user", controller.UpdateUser)
	app.Post("/api/delete_user", controller.DeleteUser)
	app.Post("/api/get_users", controller.GetCompanyUsers)
	app.Get("/api/user", controller.User)
	app.Post("/api/logout", controller.Logout)
	app.Post("/api/create_ticket", controller.CreateTicket)
	
}
