package repository

import "gorm.io/gorm"

type Repositories struct {
	UserRepo     *userRepository
	PlanningRepo *planningRepository
	ExpenseRepo  *expenseRepository
}

func InitRepositories(db *gorm.DB) *Repositories {
	userRepo := NewUserRepository(db)
	planningRepo := NewPlanningRepository(db)
	expenseRepo := NewExpenseRepository(db)
	return &Repositories{
		UserRepo:     userRepo,
		PlanningRepo: planningRepo,
		ExpenseRepo:  expenseRepo,
	}
}
