package controller

import (
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/muadzmo/go-fin-planning/models"
	"github.com/muadzmo/go-fin-planning/repository"
)

type transController struct {
	repository repository.TransRepository
	income     repository.IncomeRepository
	expense    repository.ExpenseRepository
	planning   repository.PlanningRepository
	validate   *validator.Validate
}

func NewTransController(
	repository repository.TransRepository,
	incomeRepo repository.IncomeRepository,
	expenseRepo repository.ExpenseRepository,
	planningRepo repository.PlanningRepository) *transController {
	validate := validator.New()
	return &transController{repository, incomeRepo, expenseRepo, planningRepo, validate}
}

func (t *transController) Create(c *fiber.Ctx) error {
	var data models.Transaction
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Create " + err.Error(),
		})
	}

	data.CreatedBy = c.Locals("email").(string)

	data, err := t.repository.CreateTrans(data)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "planning 77: " + err.Error(),
		})
	}

	return c.JSON(data)
}

func (t *transController) GetAll(c *fiber.Ctx) error {
	data, err := t.repository.FindAllTrans()
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(data)
}

func (t *transController) GetById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "GetById 90: " + err.Error(),
		})
	}

	data, err := t.repository.FindTransById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "GetById 97: " + err.Error(),
		})
	}

	return c.JSON(data)
}

func (t *transController) Save(c *fiber.Ctx) error {
	var data models.Transaction
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	data.ModifiedBy = c.Locals("email").(string)
	data.ModifiedAt = time.Now()

	err := t.validate.Struct(data)
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

	_, err = t.repository.FindTransById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "transaction" + err.Error(),
		})
	}

	_, err = t.planning.FindPlanningById(data.PlanId)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "plan" + err.Error(),
		})
	}

	data.Id = uint(id)
	data, err = t.repository.SaveTrans(data)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(data)
}

func (t *transController) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	_, err = t.repository.FindTransById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = t.repository.DeleteTrans(uint(id))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Delete Successfully",
	})
}
