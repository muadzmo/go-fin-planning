package models

import (
	"time"
)

type User struct {
	Id         uint
	Name       string    `json:"name" gorm:"type:varchar(255);"`
	Email      string    `json:"email" gorm:"type:varchar(255);unique"`
	Password   []byte    `json:"-"`
	createdAt  time.Time `json:"-"`
	modifiedAt time.Time `json:"-"`
}
