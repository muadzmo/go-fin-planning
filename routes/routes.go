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
	v1.Get("/user", controller.User.Get)

	v1.Get("/sofmaster/list", controller.Income.GetAllMaster)
	v1.Post("/sofmaster/add", controller.Income.CreateMaster)
	v1.Post("/sofmaster/detail/:code", controller.Income.SaveMaster)
	v1.Get("/sofmaster/detail/:code", controller.Income.GetMaster)

	v1.Get("/expensemaster/list", controller.Expense.GetAllMaster)
	v1.Post("/expensemaster/add", controller.Expense.CreateMaster)
	v1.Post("/expensemaster/detail/:code", controller.Expense.SaveMaster)
	v1.Get("/expensemaster/detail/:code", controller.Expense.GetMaster)
	v1.Delete("/expensemaster/detail/:code", controller.Expense.DeleteMaster)

	v1.Get("/planning/list", controller.Planning.GetAll)
	v1.Post("/planning/add", controller.Planning.Create)
	v1.Post("/planning/detail/:id", controller.Planning.Save)
	v1.Get("/planning/detail/:id", controller.Planning.GetById)

	v1.Get("/transaction/list", controller.Transaction.GetAll)
	v1.Post("/transaction/add", controller.Transaction.Create)
	v1.Get("/transaction/detail/:id", controller.Transaction.GetById)
	v1.Post("/transaction/detail/:id", controller.Transaction.Save)
}
