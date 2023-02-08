package controller

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/muadzmo/go-fin-planning/models"
	"github.com/muadzmo/go-fin-planning/repository"
)

type middlewareController struct {
	income   repository.IncomeRepository
	expense  repository.ExpenseRepository
	trans    repository.TransRepository
	validate *validator.Validate
}

func NewMiddlewareController(incomeRepo repository.IncomeRepository, expenseRepo repository.ExpenseRepository, transRepo repository.TransRepository) *middlewareController {
	return &middlewareController{incomeRepo, expenseRepo, transRepo, validator.New()}
}

func (m *middlewareController) ParsingPlanning(c *fiber.Ctx) error {
	var data models.Planning
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err := m.validate.Struct(data)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "planning 51: " + err.Error(),
		})
	}

	data.ModifiedBy = c.Locals("email").(string)
	data.ModifiedAt = time.Now()

	c.Locals("data", data)
	c.Locals("tipe", data.PlanType)
	c.Locals("code", data.PlanCode)
	return c.Next()
}

func (m *middlewareController) ParsingTransaction(c *fiber.Ctx) error {
	var data models.Transaction
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err := m.validate.Struct(data)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "planning 51: " + err.Error(),
		})
	}

	data.ModifiedBy = c.Locals("email").(string)
	data.ModifiedAt = time.Now()

	c.Locals("data", data)
	c.Locals("tipe", data.TransType)
	c.Locals("code", data.TransCode)
	return c.Next()
}

func (m *middlewareController) ValidateInput(c *fiber.Ctx) error {
	var err error

	tipe := c.Locals("tipe").(string)
	code := c.Locals("code").(string)

	if tipe != "expense" && tipe != "income" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Plan type is not found",
		})
	}

	if tipe == "income" {
		_, err = m.income.FindIncomeMasterByCode(code)
	} else {
		_, err = m.expense.FindExpenseMasterByCode(code)
	}
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Master code is not found",
		})
	}

	return c.Next()
}
