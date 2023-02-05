package repository

import (
	"github.com/muadzmo/go-fin-planning/models"
	"gorm.io/gorm"
)

type transRepository struct {
	DB *gorm.DB
}

type TransRepository interface {
	FindTransById(id uint) (models.Transaction, error)
	CreateTrans(data models.Transaction) (models.Transaction, error)
	FindAllTrans() ([]models.Transaction, error)
	SaveTrans(data models.Transaction) (models.Transaction, error)
}

func NewTransRepository(db *gorm.DB) *transRepository {
	return &transRepository{DB: db}
}

func (t *transRepository) FindTransById(id uint) (models.Transaction, error) {
	var trans models.Transaction
	err := t.DB.First(&trans)
	return trans, err.Error
}

func (t *transRepository) CreateTrans(data models.Transaction) (models.Transaction, error) {
	err := t.DB.Create(&data)
	return data, err.Error
}

func (t *transRepository) FindAllTrans() ([]models.Transaction, error) {
	var data []models.Transaction
	err := t.DB.Find(&data)
	return data, err.Error
}

func (t *transRepository) SaveTrans(data models.Transaction) (models.Transaction, error) {
	err := t.DB.Where("id = ?", data.Id).Updates(&data)
	return data, err.Error
}

func (t *transRepository) DeleteTrans(data models.Transaction) error {
	err := t.DB.Delete(&data)
	return err.Error
}
