package repository

import (
	"github.com/muadzmo/go-fin-planning/models"
	"gorm.io/gorm"
)

type balanceRepository struct {
	DB *gorm.DB
}

type BalanceRepository interface {
	Create(balance models.Balance) (models.Balance, error)
	FindAll() ([]models.Balance, error)
	FindByCode(code string) (models.Balance, error)
	Save(balance models.Balance) (models.Balance, error)
	Delete(balance models.Balance) error
}

func NewBalanceRepository(db *gorm.DB) *balanceRepository {
	return &balanceRepository{
		DB: db,
	}
}

func (b *balanceRepository) Create(balance models.Balance) (models.Balance, error) {
	err := b.DB.Create(&balance)
	return balance, err.Error
}

func (b *balanceRepository) FindAll() ([]models.Balance, error) {
	var balances []models.Balance
	err := b.DB.Find(&balances)
	return balances, err.Error
}

func (b *balanceRepository) FindByCode(code string) (models.Balance, error) {
	var balance models.Balance
	err := b.DB.Where("code = ?", code).First(&balance)
	return balance, err.Error
}

func (b *balanceRepository) Save(balance models.Balance) (models.Balance, error) {
	err := b.DB.Save(&balance)
	return balance, err.Error
}

func (b *balanceRepository) Delete(balance models.Balance) error {
	err := b.DB.Where("code = ?", balance.Code).Delete(&balance)
	return err.Error
}
