package repository

import (
	"github.com/muadzmo/go-fin-planning/models"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	FindUserById(id int) (models.User, error)
	FindUserByEmail(email string) (models.User, error)
	CreateUser(models.User) (models.User, error)
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		DB: db,
	}
}

func (u userRepository) FindUserById(id int) (models.User, error) {
	var user models.User
	err := u.DB.Where("id = ?", id).First(&user)
	return user, err.Error
}

func (u *userRepository) FindUserByEmail(email string) (models.User, error) {
	var user models.User
	err := u.DB.Where("email = ?", email).First(&user)
	return user, err.Error
}

func (u *userRepository) CreateUser(user models.User) (models.User, error) {
	err := u.DB.Create(&user)
	return user, err.Error
}
