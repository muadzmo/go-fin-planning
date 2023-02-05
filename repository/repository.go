package repository

import "gorm.io/gorm"

type Repositories struct {
	UserRepo     *userRepository
	IncomeRepo   *incomeRepository
	ExpenseRepo  *expenseRepository
	PlanningRepo *planningRepository
	TransRepo    *transRepository
}

func InitRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepo:     NewUserRepository(db),
		IncomeRepo:   NewIncomeRepository(db),
		ExpenseRepo:  NewExpenseRepository(db),
		PlanningRepo: NewPlanningRepository(db),
		TransRepo:    NewTransRepository(db),
	}
}
