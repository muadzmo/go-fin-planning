package models

import "time"

type PlanningHeader struct {
	PlanDate time.Time `json:"plan_date" gorm:"not null" validate:"required"`
	Amount   float32   `json:"plan_amount" gorm:"not null" validate:"required"`
}

type Planning struct {
	Id uint `json:"id" gorm:"primaryKey, autoIncrement"`
	PlanningHeader
	BalanceCode string    `json:"balance_code" gorm:"type:varchar(50);not null" validate:"required"`
	CreatedAt   time.Time `json:"-"`
	CreatedBy   string    `json:"-" gorm:"type:varchar(255);"`
	ModifiedAt  time.Time `json:"-"`
	ModifiedBy  string    `json:"-" gorm:"type:varchar(255);"`
}

type PlanningDetail struct {
	PlanningHeader
	BalanceHeader
}
