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

	api := app.Group("/api")
	api.Post("/register", controller.Auth.Register)
	api.Post("/login", controller.Auth.Login)
	api.Post("/logout", controller.Auth.Logout)

	v1 := api.Group("/v1", controller.Auth.CheckLoggedIn)
	v1.Get("/user", controller.User.GetUser)
	v1.Get("/sofmaster/list", controller.Income.GetAllMaster)
	v1.Post("/sofmaster/add", controller.Income.CreateMaster)
	v1.Post("/sofmaster/:code", controller.Income.SaveMaster)
	v1.Get("/sofmaster/:code", controller.Income.GetMaster)

	v1.Get("/expensemaster/list", controller.Expense.GetAllMaster)
	v1.Post("/expensemaster/add", controller.Expense.CreateMaster)
	v1.Post("/expensemaster/:code", controller.Expense.SaveMaster)
	v1.Get("/expensemaster/:code", controller.Expense.GetMaster)
	v1.Delete("/expensemaster/:code", controller.Expense.DeleteMaster)

}
