package models

import "time"

type Balance struct {
	BalanceHeader
	CreatedAt  time.Time `json:"-"`
	CreatedBy  string    `json:"-" gorm:"varchar(255);"`
	ModifiedAt time.Time `json:"-"`
	ModifiedBy string    `json:"-" gorm:"varchar(255);"`
}

type BalanceHeader struct {
	Code     string `json:"code" gorm:"type:varchar(50);primaryKey,index:code,unique" validate:"required,min=3,alphanum"`
	Name     string `json:"name" gorm:"type:varchar(255);" validate:"required,min=3"`
	Periodic string `json:"periodic" gorm:"type:varchar(10);"`
	Type     string `json:"type" gorm:"type:varchar(10);not null" validate:"required,alphanum"`
}
