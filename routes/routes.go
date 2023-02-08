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

	parsingPlanning := controller.Parsing.ParsingPlanning
	parsingTransaction := controller.Parsing.ParsingTransaction
	validateInput := controller.Validate.PlanningTransaction
	v1.Get("/planning/list", controller.Planning.GetAll)
	v1.Post("/planning/add", parsingPlanning, validateInput, controller.Planning.Create)
	v1.Post("/planning/detail/:id", parsingPlanning, validateInput, controller.Planning.Save)
	v1.Get("/planning/detail/:id", controller.Planning.GetById)

	v1.Get("/transaction/list", controller.Transaction.GetAll)
	v1.Post("/transaction/add", parsingTransaction, validateInput, controller.Transaction.Create)
	v1.Get("/transaction/detail/:id", controller.Transaction.GetById)
	v1.Post("/transaction/detail/:id", parsingTransaction, validateInput, controller.Transaction.Save)
	v1.Delete("/transaction/detail/:id", controller.Transaction.Delete)

	v1.Get("/balance/list", controller.Balance.GetAllMaster)
	v1.Post("/balance/add", controller.Parsing.ParsingBalance, controller.Balance.Add)
	v1.Get("/balance/detail/:code", controller.Balance.GetMaster)
	v1.Post("/balance/detail/:code", controller.Balance.GetMaster, controller.Balance.Save)
	v1.Delete("/balance/detail/:code", controller.Balance.Delete)
}
