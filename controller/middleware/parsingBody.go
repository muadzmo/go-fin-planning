package middleware

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/muadzmo/go-fin-planning/models"
	"github.com/muadzmo/go-fin-planning/repository"
)

type ParsingController struct {
	balance  repository.BalanceRepository
	validate *validator.Validate
}

func NewParsingController(balance repository.BalanceRepository) *ParsingController {
	return &ParsingController{balance, validator.New()}
}

func (m *ParsingController) ParsingPlanning(c *fiber.Ctx) error {
	var data models.Planning
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err := m.validate.Struct(data)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "parsing planning: " + err.Error(),
		})
	}

	data.ModifiedBy = c.Locals("email").(string)
	data.ModifiedAt = time.Now()

	c.Locals("data", data)
	c.Locals("code", data.BalanceCode)
	return c.Next()
}

func (m *ParsingController) ParsingTransaction(c *fiber.Ctx) error {
	var data models.Transaction
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
	}

	err := m.validate.Struct(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	data.ModifiedBy = c.Locals("email").(string)
	data.ModifiedAt = time.Now()

	c.Locals("data", data)
	c.Locals("code", data.BalanceCode)

	return c.Next()
}

func (m *ParsingController) ParsingBalance(c *fiber.Ctx) error {
	var data models.Balance
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err := m.validate.Struct(data)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "parsing balance: " + err.Error(),
		})
	}

	data.ModifiedAt = time.Now()
	data.ModifiedBy = c.Locals("email").(string)

	c.Locals("data", data)
	return c.Next()
}
