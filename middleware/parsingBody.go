package middleware

import (
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/muadzmo/go-fin-planning/models"
)

type middlewareController struct {
	validate *validator.Validate
}

func NewMiddlewareController() *middlewareController {
	return &middlewareController{validate: validator.New()}
}

func (m *middlewareController) ParsingPlanning(c *fiber.Ctx) error {
	var data models.Planning
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	data.ModifiedAt = time.Now()
	data.ModifiedBy = c.Locals("email").(string)

	c.Locals("data", data)
	return c.Next()
}

func (m *middlewareController) ParsingTransaction(c *fiber.Ctx) error {
	var data models.Transaction
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	data.ModifiedAt = time.Now()
	data.ModifiedBy = c.Locals("email").(string)

	c.Locals("data", data)
	return c.Next()
}
