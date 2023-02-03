package repository

import (
	"github.com/muadzmo/go-fin-planning/models"
	"gorm.io/gorm"
)

type expenseRepository struct {
	DB *gorm.DB
}

type ExpenseRepository interface {
	CreateExpenseMaster(expense models.MasterExpense) (models.MasterExpense, error)
	FindAllExpenseMaster() ([]models.MasterExpense, error)
	FindExpenseMasterByCode(expense models.MasterExpense, code string) (models.MasterExpense, error)
	SaveExpenseMaster(expense models.MasterExpense) (models.MasterExpense, error)
	DeleteExpenseMaster(expense models.MasterExpense, code string) error
}

func NewExpenseRepository(db *gorm.DB) *expenseRepository {
	return &expenseRepository{
		DB: db,
	}
}

func (e *expenseRepository) CreateExpenseMaster(expense models.MasterExpense) (models.MasterExpense, error) {
	err := e.DB.Create(&expense)
	return expense, err.Error
}

func (e *expenseRepository) FindAllExpenseMaster() ([]models.MasterExpense, error) {
	var expenses []models.MasterExpense
	err := e.DB.Find(&expenses)
	return expenses, err.Error
}

func (e *expenseRepository) FindExpenseMasterByCode(expense models.MasterExpense, code string) (models.MasterExpense, error) {
	err := e.DB.Where("code = ?", code).First(&expense)
	return expense, err.Error
}

func (e *expenseRepository) SaveExpenseMaster(expense models.MasterExpense) (models.MasterExpense, error) {
	err := e.DB.Save(&expense)
	return expense, err.Error
}

func (e *expenseRepository) DeleteExpenseMaster(expense models.MasterExpense, code string) error {
	err := e.DB.Delete("code = ?", code)
	return err.Error
}
