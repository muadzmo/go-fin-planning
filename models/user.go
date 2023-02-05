package models

import (
	"time"
)

type User struct {
	Id         uint
	Name       string    `json:"name"`
	Email      string    `json:"email" gorm:"unique"`
	Password   []byte    `json:"-"`
	createdAt  time.Time `json:"-"`
	modifiedAt time.Time `json:"-"`
}
