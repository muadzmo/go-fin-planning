package controller

import (
	"strconv"

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
	data := c.Locals("data").(models.Planning)
	data.CreatedBy = c.Locals("email").(string)

	data, err := p.repository.CreatePlanning(data)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(data)
}

func (p *planningController) Save(c *fiber.Ctx) error {
	data := c.Locals("data").(models.Planning)

	// get id by url param
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// get previous data, if not exist return false
	_, err = p.repository.FindPlanningById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// save data
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

	data, err := p.repository.FindPlanningById(uint(id))
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
