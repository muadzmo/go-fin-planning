package repository

import (
	"github.com/muadzmo/go-fin-planning/models"
	"gorm.io/gorm"
)

type transRepository struct {
	DB *gorm.DB
}

type TransRepository interface {
	FindTransById(id uint) (models.TransactionDetail, error)
	CreateTrans(data models.Transaction) (models.Transaction, error)
	FindAllTrans() ([]models.Transaction, error)
	SaveTrans(data models.Transaction) (models.Transaction, error)
	DeleteTrans(id uint) error
}

func NewTransRepository(db *gorm.DB) *transRepository {
	return &transRepository{DB: db}
}

func (t *transRepository) FindTransById(id uint) (models.TransactionDetail, error) {
	var transDetail models.TransactionDetail
	err := t.DB.Model(&models.Transaction{}).
		Select("transactions.id, trans_date, transactions.amount, code, b.name, b.periodic, b.type, plan_id, p.plan_date PlanDate, p.amount").
		Joins("join balances b on transactions.balance_code = b.code").
		Joins("left join plannings p on p.id = transactions.plan_id").
		Where("transactions.id = ?", id).
		First(&transDetail)
	return transDetail, err.Error
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

func (t *transRepository) DeleteTrans(id uint) error {
	var data models.Transaction
	err := t.DB.Where("id = ?", id).Delete(&data)
	return err.Error
}
