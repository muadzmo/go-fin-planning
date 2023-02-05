package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/muadzmo/go-fin-planning/repository"
)

type userController struct {
	repository repository.UserRepository
}

func NewUserController(repository repository.UserRepository) *userController {
	return &userController{
		repository,
	}
}

func (u *userController) Get(c *fiber.Ctx) error {
	str, ok := c.Locals("id").(string)
	if !ok {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "id is not string",
		})
	}

	userId, err := strconv.Atoi(str)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user, err := u.repository.FindUserById(userId)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(user)
}
