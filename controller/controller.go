package controller

import (
	expensetype "github.com/muadzmo/go-fin-planning/controller/expenseType"
	incometype "github.com/muadzmo/go-fin-planning/controller/incomeType"
	"github.com/muadzmo/go-fin-planning/repository"
)

type Controllers struct {
	Auth    *authController
	User    *userController
	Income  *incometype.IncomeTypeController
	Expense *expensetype.ExpenseTypeController
}

func InitControllers(repo repository.Repositories) *Controllers {
	return &Controllers{
		Auth:    NewAuthController(repo.UserRepo),
		User:    NewUserController(repo.UserRepo),
		Income:  incometype.NewSourceOfFundMasterController(repo.PlanningRepo),
		Expense: expensetype.NewExpenseTypeMasterController(repo.ExpenseRepo),
	}
}
