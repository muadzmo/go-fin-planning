package repository

import (
	"github.com/muadzmo/go-fin-planning/models"
	"gorm.io/gorm"
)

type incomeRepository struct {
	DB *gorm.DB
}

type IncomeRepository interface {
	CreateIncomeMaster(income models.MasterIncome) (models.MasterIncome, error)
	SaveIncomeMaster(income models.MasterIncome) (models.MasterIncome, error)
	FindIncomeMasterByCode(code string) (models.MasterIncome, error)
	FindAllIncomeMaster() ([]models.MasterIncome, error)
}

func NewIncomeRepository(db *gorm.DB) *incomeRepository {
	return &incomeRepository{
		DB: db,
	}
}

func (i *incomeRepository) CreateIncomeMaster(income models.MasterIncome) (models.MasterIncome, error) {
	err := i.DB.Create(&income)
	return income, err.Error
}

func (i *incomeRepository) SaveIncomeMaster(income models.MasterIncome) (models.MasterIncome, error) {
	err := i.DB.Save(&income)
	return income, err.Error
}

func (i *incomeRepository) FindIncomeMasterByCode(code string) (models.MasterIncome, error) {
	var income models.MasterIncome
	err := i.DB.Where("code = ?", code).First(&income)
	return income, err.Error
}

func (i *incomeRepository) FindAllIncomeMaster() ([]models.MasterIncome, error) {
	var sourcesOfFund []models.MasterIncome
	err := i.DB.Find(&sourcesOfFund)
	return sourcesOfFund, err.Error
}
