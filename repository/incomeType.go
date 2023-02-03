package repository

import (
	"github.com/muadzmo/go-fin-planning/models"
	"gorm.io/gorm"
)

type planningRepository struct {
	DB *gorm.DB
}

type PlanningRepository interface {
	CreateIncomeMaster(sourceOfFund models.MasterIncome) (models.MasterIncome, error)
	SaveIncomeMaster(sourceOfFund models.MasterIncome) (models.MasterIncome, error)
	FindIncomeMasterByCode(sourceOfFund models.MasterIncome, code string) (models.MasterIncome, error)
	FindAllIncomeMaster() ([]models.MasterIncome, error)
}

func NewPlanningRepository(db *gorm.DB) *planningRepository {
	return &planningRepository{
		DB: db,
	}
}

func (p *planningRepository) CreateIncomeMaster(sourceOfFund models.MasterIncome) (models.MasterIncome, error) {
	err := p.DB.Create(&sourceOfFund)
	return sourceOfFund, err.Error
}

func (p *planningRepository) SaveIncomeMaster(sourceOfFund models.MasterIncome) (models.MasterIncome, error) {
	err := p.DB.Save(&sourceOfFund)
	return sourceOfFund, err.Error
}

func (p *planningRepository) FindIncomeMasterByCode(sourceOfFund models.MasterIncome, code string) (models.MasterIncome, error) {
	err := p.DB.Where("code = ?", code).First(&sourceOfFund)
	return sourceOfFund, err.Error
}

func (p *planningRepository) FindAllIncomeMaster() ([]models.MasterIncome, error) {
	var sourcesOfFund []models.MasterIncome
	err := p.DB.Find(&sourcesOfFund)
	return sourcesOfFund, err.Error
}
