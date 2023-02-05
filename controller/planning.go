package controller

import (
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/muadzmo/go-fin-planning/models"
	"github.com/muadzmo/go-fin-planning/repository"
)

type planningController struct {
	repository  repository.PlanningRepository
	income      repository.IncomeRepository
	expense     repository.ExpenseRepository
	transaction repository.TransRepository
	validate    *validator.Validate
}

func NewPlanningController(
	repository repository.PlanningRepository,
	incomeRepo repository.IncomeRepository,
	expenseRepo repository.ExpenseRepository,
	transRepo repository.TransRepository) *planningController {
	validate := validator.New()
	return &planningController{repository, incomeRepo, expenseRepo, transRepo, validate}
}

func (p *planningController) GetAll(c *fiber.Ctx) error {
	data, err := p.repository.FindAllPlanning()
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(data)
}

func (p *planningController) Create(c *fiber.Ctx) error {
	var data models.Planning
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Create " + err.Error(),
		})
	}

	data.CreatedBy = c.Locals("email").(string)

	err := p.validate.Struct(data)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "planning 51: " + err.Error(),
		})
	}

	if data.PlanType != "expense" && data.PlanType != "income" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Plan type is not found",
		})
	}

	if data.PlanType == "income" {
		var income models.MasterIncome
		_, err = p.income.FindIncomeMasterByCode(income, data.PlanCode)
	} else {
		var expense models.MasterExpense
		_, err = p.expense.FindExpenseMasterByCode(expense, data.PlanCode)
	}
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Master code is not found",
		})
	}

	data, err = p.repository.CreatePlanning(data)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "planning 77: " + err.Error(),
		})
	}

	return c.JSON(data)
}

func (p *planningController) Save(c *fiber.Ctx) error {
	var data models.Planning
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	data.ModifiedBy = c.Locals("email").(string)
	data.ModifiedAt = time.Now()

	err := p.validate.Struct(data)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var origin models.Planning
	origin, err = p.repository.FindPlanningById(origin, uint(id))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	data.Id = uint(id)
	data, err = p.repository.SavePlanning(data)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(data)
}

func (p *planningController) GetById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "GetById " + err.Error(),
		})
	}

	var data models.Planning
	data, err = p.repository.FindPlanningById(data, uint(id))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var detail models.PlanningDetail
	detail.Amount = data.Amount
	detail.Id = data.Id
	detail.PlanCode = data.PlanCode
	detail.PlanType = data.PlanType
	detail.PlanDate = data.PlanDate
	detail.TransId = data.TransId

	if detail.TransId != 0 {
		transDetail, err := p.transaction.FindTransById(detail.TransId)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		detail.TransDate = transDetail.TransDate
		detail.TransType = transDetail.TransType
	}

	return c.JSON(detail)
}
