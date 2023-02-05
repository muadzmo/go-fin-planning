package income

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/muadzmo/go-fin-planning/models"
	"github.com/muadzmo/go-fin-planning/repository"
)

type IncomeTypeController struct {
	repository repository.IncomeRepository
	validate   *validator.Validate
}

func NewSourceOfFundMasterController(repository repository.IncomeRepository) *IncomeTypeController {
	validate := validator.New()
	return &IncomeTypeController{
		repository,
		validate,
	}
}

func (s *IncomeTypeController) CreateMaster(c *fiber.Ctx) error {
	var data = new(models.MasterIncome)

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	data.CreatedBy = c.Locals("email").(string)

	err := s.validate.Struct(data)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	sofMaster, err := s.repository.CreateIncomeMaster(*data)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(sofMaster)
}

func (s *IncomeTypeController) SaveMaster(c *fiber.Ctx) error {
	var data models.MasterIncome
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	dataMaster, err := s.repository.FindIncomeMasterByCode(data, c.Params("code"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	dataMaster.Name = data.Name
	dataMaster.ModifiedBy = c.Locals("email").(string)
	dataMaster.ModifiedAt = time.Now()

	err = s.validate.Struct(dataMaster)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	sofMaster, err := s.repository.SaveIncomeMaster(dataMaster)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(sofMaster)
}

func (s *IncomeTypeController) GetAllMaster(c *fiber.Ctx) error {
	dataMaster, err := s.repository.FindAllIncomeMaster()
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(dataMaster)
}

func (s *IncomeTypeController) GetMaster(c *fiber.Ctx) error {
	var data models.MasterIncome
	data, err := s.repository.FindIncomeMasterByCode(data, c.Params("code"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(data)
}
