package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/muadzmo/go-fin-planning/repository"
)

type ValidateController struct {
	balance  repository.BalanceRepository
	validate *validator.Validate
}

func NewValidateController(balance repository.BalanceRepository) *ValidateController {
	return &ValidateController{balance, validator.New()}
}

func (v *ValidateController) PlanningTransaction(c *fiber.Ctx) error {
	code := c.Locals("code").(string)

	_, err := v.balance.FindByCode(code)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "balance " + err.Error(),
		})
	}

	return c.Next()
}
