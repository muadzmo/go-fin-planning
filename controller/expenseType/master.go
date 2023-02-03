package expensetype

import (
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/muadzmo/go-fin-planning/models"
	"github.com/muadzmo/go-fin-planning/repository"
)

type ExpenseTypeController struct {
	repository repository.ExpenseRepository
	validate   *validator.Validate
}

func NewExpenseTypeMasterController(repository repository.ExpenseRepository) *ExpenseTypeController {
	validate := validator.New()
	return &ExpenseTypeController{
		repository,
		validate,
	}
}

func (e *ExpenseTypeController) CreateMaster(c *fiber.Ctx) error {
	var data models.MasterExpense
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	data.CreatedBy = c.Locals("email").(string)

	data, err := e.repository.CreateExpenseMaster(data)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(data)
}

func (e *ExpenseTypeController) GetAllMaster(c *fiber.Ctx) error {
	data, err := e.repository.FindAllExpenseMaster()
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(data)
}

func (e *ExpenseTypeController) SaveMaster(c *fiber.Ctx) error {
	var data models.MasterExpense
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	dataMaster, err := e.repository.FindExpenseMasterByCode(data, c.Params("code"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	dataMaster.Name = data.Name
	dataMaster.ModifiedBy = c.Locals("email").(string)
	dataMaster.ModifiedAt = time.Now()

	err = e.validate.Struct(dataMaster)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	sofMaster, err := e.repository.SaveExpenseMaster(dataMaster)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(sofMaster)
}

func (e *ExpenseTypeController) GetMaster(c *fiber.Ctx) error {
	var data models.MasterExpense
	data, err := e.repository.FindExpenseMasterByCode(data, c.Params("code"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(data)
}

func (e *ExpenseTypeController) DeleteMaster(c *fiber.Ctx) error {
	var data models.MasterExpense
	data, err := e.repository.FindExpenseMasterByCode(data, c.Params("code"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(data)
}
