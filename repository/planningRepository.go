package repository

import (
	"github.com/muadzmo/go-fin-planning/models"
	"gorm.io/gorm"
)

type planningRepository struct {
	DB *gorm.DB
}

type PlanningRepository interface {
	CreateSourceOfFoundMaster(sourceOfFund models.SourceOfFundMaster) (models.SourceOfFundMaster, error)
	SaveSourceOfFoundMaster(sourceOfFund models.SourceOfFundMaster) (models.SourceOfFundMaster, error)
	GetSourceOfFoundMaster(sourceOfFund models.SourceOfFundMaster, code string) (models.SourceOfFundMaster, error)
}

func NewPlanningRepository(db *gorm.DB) *planningRepository {
	return &planningRepository{
		DB: db,
	}
}

func (p *planningRepository) CreateSourceOfFoundMaster(sourceOfFund models.SourceOfFundMaster) (models.SourceOfFundMaster, error) {
	err := p.DB.Create(&sourceOfFund)
	return sourceOfFund, err.Error
}

func (p *planningRepository) SaveSourceOfFoundMaster(sourceOfFund models.SourceOfFundMaster) (models.SourceOfFundMaster, error) {
	err := p.DB.Save(&sourceOfFund)
	return sourceOfFund, err.Error
}

func (p *planningRepository) GetSourceOfFoundMaster(sourceOfFund models.SourceOfFundMaster, code string) (models.SourceOfFundMaster, error) {
	err := p.DB.Where("code = ?", code).First(&sourceOfFund)
	return sourceOfFund, err.Error
}
