package controller

import (
	sourceoffund "github.com/muadzmo/go-fin-planning/controller/sourceOfFund"
	"github.com/muadzmo/go-fin-planning/repository"
)

type Controllers struct {
	AuthC      *authController
	UserC      *userController
	SoFMasterC *sourceoffund.SofController
}

func InitControllers(repo repository.Repositories) *Controllers {
	return &Controllers{
		AuthC:      NewAuthController(repo.UserRepo),
		UserC:      NewUserController(repo.UserRepo),
		SoFMasterC: sourceoffund.NewSourceOfFundMasterController(repo.PlanningRepo),
	}
}
