package models

import (
	"time"
)

type SourceOfFundMaster struct {
	Code       string `json:"code" gorm:"primaryKey" validate:"required,min=3,alphanum"`
	Name       string `json:"name" validate:"required,min=3"`
	CreatedAt  time.Time
	CreatedBy  string
	ModifiedAt time.Time
	ModifiedBy string
}
