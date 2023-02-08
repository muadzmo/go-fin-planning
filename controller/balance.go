package controller

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/muadzmo/go-fin-planning/models"
	"github.com/muadzmo/go-fin-planning/repository"
)

type balanceController struct {
	repository repository.BalanceRepository
	validate   *validator.Validate
}

func NewBalanceController(repository repository.BalanceRepository) *balanceController {
	validate := validator.New()
	return &balanceController{
		repository,
		validate,
	}
}

func (b *balanceController) Add(c *fiber.Ctx) error {
	data := c.Locals("data").(models.Balance)
	data.CreatedBy = c.Locals("email").(string)

	master, err := b.repository.Create(data)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(master)
}

func (b *balanceController) Save(c *fiber.Ctx) error {
	data := c.Locals("data").(models.Balance)

	master, err := b.repository.FindByCode(c.Params("code"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// to do : if balance is already used by (transaction or planning) and (what to edit is periodic or type), return error.

	master.Name = data.Name
	master.Periodic = data.Periodic
	master.Type = data.Type
	master.ModifiedBy = c.Locals("email").(string)
	master.ModifiedAt = time.Now()

	err = b.validate.Struct(master)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	master, err = b.repository.Save(master)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(master)
}

func (b *balanceController) GetAllMaster(c *fiber.Ctx) error {
	dataMaster, err := b.repository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(dataMaster)
}

func (b *balanceController) GetMaster(c *fiber.Ctx) error {
	data, err := b.repository.FindByCode(c.Params("code"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// return c.Next()

	return c.JSON(data)
}

func (b *balanceController) Delete(c *fiber.Ctx) error {
	data, err := b.repository.FindByCode(c.Params("code"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// to do : if balance is already used by (transaction or planning) return error.

	err = b.repository.Delete(data)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully delete balance",
	})
}
