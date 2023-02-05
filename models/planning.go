package models

import "time"

type PlanningHeader struct {
	Id       uint      `json:"id" gorm:"primaryKey, autoIncrement"`
	PlanDate time.Time `json:"plan_date" gorm:"not null" validate:"required"`
	PlanType string    `json:"plan_type" gorm:"not null" validate:"required"`
	PlanCode string    `json:"plan_code" gorm:"not null" validate:"required"`
	Amount   float32   `json:"amount" gorm:"not null" validate:"required"`
	TransId  uint      `json:"trans_id"`
}

type Planning struct {
	PlanningHeader
	CreatedAt  time.Time `json:"-"`
	CreatedBy  string    `json:"-"`
	ModifiedAt time.Time `json:"-"`
	ModifiedBy string    `json:"-"`
}

type PlanningDetail struct {
	PlanningHeader
	TransactionDetail
}
