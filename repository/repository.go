package repository

import "gorm.io/gorm"

type Repositories struct {
	UserRepo     *userRepository
	PlanningRepo *planningRepository
}

func InitRepositories(db *gorm.DB) *Repositories {
	userRepo := NewUserRepository(db)
	planningRepo := NewPlanningRepository(db)
	return &Repositories{
		UserRepo:     userRepo,
		PlanningRepo: planningRepo,
	}
}
