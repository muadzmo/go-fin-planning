package controller

import (
	expense "github.com/muadzmo/go-fin-planning/controller/expense"
	income "github.com/muadzmo/go-fin-planning/controller/income"
	"github.com/muadzmo/go-fin-planning/controller/middleware"
	"github.com/muadzmo/go-fin-planning/repository"
)

type Controllers struct {
	Auth        *authController
	User        *userController
	Income      *income.IncomeTypeController
	Expense     *expense.ExpenseTypeController
	Planning    *planningController
	Transaction *transController
	Balance     *balanceController
	Parsing     *middleware.ParsingController
	Validate    *middleware.ValidateController
}

func InitControllers(repo repository.Repositories) *Controllers {
	return &Controllers{
		Auth:        NewAuthController(repo.UserRepo),
		User:        NewUserController(repo.UserRepo),
		Income:      income.NewSourceOfFundMasterController(repo.IncomeRepo),
		Expense:     expense.NewExpenseTypeMasterController(repo.ExpenseRepo),
		Planning:    NewPlanningController(repo.PlanningRepo, repo.IncomeRepo, repo.ExpenseRepo, repo.TransRepo),
		Transaction: NewTransController(repo.TransRepo, repo.IncomeRepo, repo.ExpenseRepo, repo.PlanningRepo),
		Balance:     NewBalanceController(repo.BalanceRepo),
		Parsing:     middleware.NewParsingController(repo.BalanceRepo),
		Validate:    middleware.NewValidateController(repo.BalanceRepo),
	}
}
