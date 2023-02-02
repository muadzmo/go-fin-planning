package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muadzmo/go-fin-planning/controller"
	"github.com/muadzmo/go-fin-planning/database"
	"github.com/muadzmo/go-fin-planning/repository"
)

func Setup(app *fiber.App) {
	repo := repository.InitRepositories(database.DB)
	controller := controller.InitControllers(*repo)

	app.Post("/api/register", controller.AuthC.Register)
	app.Post("/api/login", controller.AuthC.Login)
	app.Post("/api/logout", controller.AuthC.Logout)

	app.Get("/api/user", controller.AuthC.CheckLoggedIn, controller.UserC.GetUser)
	app.Post("/api/sofmaster/add", controller.AuthC.CheckLoggedIn, controller.SoFMasterC.CreateMaster)
	app.Post("/api/sofmaster/save/:code", controller.AuthC.CheckLoggedIn, controller.SoFMasterC.SaveMaster)

}
