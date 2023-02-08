package models

import "time"

type Transaction struct {
	TransactionHeader
	BalanceCode string    `json:"balance_code" gorm:"type:varchar(50);not null" validate:"required"`
	CreatedAt   time.Time `json:"-"`
	CreatedBy   string    `json:"-" gorm:"type:varchar(255);"`
	ModifiedAt  time.Time `json:"-"`
	ModifiedBy  string    `json:"-" gorm:"type:varchar(255);"`
}

type TransactionHeader struct {
	Id        uint      `json:"id" gorm:"primaryKey, autoIncrement"`
	TransDate time.Time `json:"trans_date" gorm:"not null" validate:"required"`
	Amount    float32   `json:"trans_amount" gorm:"not null" validate:"required"`
	PlanId    uint      `json:"plan_id" gorm:"not null"`
}

type TransactionDetail struct {
	TransactionHeader
	BalanceHeader
	PlanningHeader
}
