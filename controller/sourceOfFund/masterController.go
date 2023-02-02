package sourceoffund

import (
	"fmt"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/muadzmo/go-fin-planning/models"
	"github.com/muadzmo/go-fin-planning/repository"
)

type SourceOfFundMaster interface {
}

type SofController struct {
	repository repository.PlanningRepository
}

func NewSourceOfFundMasterController(repository repository.PlanningRepository) *SofController {
	return &SofController{
		repository,
	}
}

func (s *SofController) CreateMaster(c *fiber.Ctx) error {
	var data = new(models.SourceOfFundMaster)

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	data.CreatedBy = c.Locals("email").(string)

	validate := validator.New()
	err := validate.Struct(data)
	fmt.Println(err)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	sofMaster, err := s.repository.CreateSourceOfFoundMaster(*data)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(sofMaster)
}

func (s *SofController) SaveMaster(c *fiber.Ctx) error {
	var data models.SourceOfFundMaster
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	dataMaster, err := s.repository.GetSourceOfFoundMaster(data, c.Params("code"))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	dataMaster.Name = data.Name
	dataMaster.ModifiedBy = c.Locals("email").(string)
	dataMaster.ModifiedAt = time.Now()

	validate := validator.New()
	err = validate.Struct(dataMaster)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	sofMaster, err := s.repository.SaveSourceOfFoundMaster(dataMaster)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(sofMaster)
}
