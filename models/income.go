package models

import (
	"time"
)

type MasterIncome struct {
	Code       string    `json:"code" gorm:"primaryKey" validate:"required,min=3,alphanum"`
	Name       string    `json:"name" validate:"required,min=3"`
	Periodic   string    `json:"periodic"`
	CreatedAt  time.Time `json:"-"`
	CreatedBy  string    `json:"-"`
	ModifiedAt time.Time `json:"-"`
	ModifiedBy string    `json:"-"`
}
