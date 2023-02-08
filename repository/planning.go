package repository

import (
	"github.com/muadzmo/go-fin-planning/models"
	"gorm.io/gorm"
)

type planningRepository struct {
	DB *gorm.DB
}

type PlanningRepository interface {
	FindAllPlanning() ([]models.Planning, error)
	FindPlanningById(id uint) (models.Planning, error)
	CreatePlanning(data models.Planning) (models.Planning, error)
	SavePlanning(data models.Planning) (models.Planning, error)
}

func NewPlanningRepository(db *gorm.DB) *planningRepository {
	return &planningRepository{DB: db}
}

func (p *planningRepository) FindAllPlanning() ([]models.Planning, error) {
	var planning []models.Planning
	err := p.DB.Find(&planning)
	return planning, err.Error
}

func (p *planningRepository) FindPlanningById(id uint) (models.Planning, error) {
	var data models.Planning
	err := p.DB.Where("id = ?", id).First(&data)
	return data, err.Error
}

func (p *planningRepository) CreatePlanning(data models.Planning) (models.Planning, error) {
	err := p.DB.Create(&data)
	return data, err.Error
}

func (p *planningRepository) SavePlanning(data models.Planning) (models.Planning, error) {
	// err := p.DB.Save(&data)
	err := p.DB.Where("id = ?", data.Id).Updates(&data)
	// err := p.DB.Model(&data).Where("id = ?", data.Id).Updates(&data)
	return data, err.Error
}
