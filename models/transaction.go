package models

import "time"

type Transaction struct {
	Id uint `json:"id" gorm:"primaryKey, autoIncrement"`
	TransactionDetail
	TransCode  string    `json:"trans_code" gorm:"not null" validate:"required"`
	CreatedAt  time.Time `json:"-"`
	CreatedBy  string    `json:"-"`
	ModifiedAt time.Time `json:"-"`
	ModifiedBy string    `json:"-"`
}

type TransactionDetail struct {
	TransDate   time.Time `json:"trans_date" gorm:"not null" validate:"required"`
	TransType   string    `json:"trans_type" gorm:"not null" validate:"required"`
	TransAmount float32   `json:"trans_amount" gorm:"not null" validate:"required"`
	PlanId      uint      `json:"plan_id" gorm:"not null"`
}
