package controller

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/muadzmo/go-fin-planning/models"
	"github.com/muadzmo/go-fin-planning/repository"
	"golang.org/x/crypto/bcrypt"
)

type authController struct {
	repository repository.UserRepository
}

func NewAuthController(repository repository.UserRepository) *authController {
	return &authController{repository}
}

func (a *authController) Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	user, err := a.repository.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(user)
}

func (a *authController) Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user, err := a.repository.FindUserByEmail(data["email"])
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        strconv.Itoa(int(user.Id)),
		Issuer:    user.Email,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	SecretKey := os.Getenv("secret_key")
	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
		"name":    user.Name,
	})
}

func (a *authController) Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func (a *authController) CheckLoggedIn(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "missing token",
		})
	}

	secretKey := os.Getenv("secret_key")
	if secretKey == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "secret key not set",
		})
	}

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid token",
		})
	}

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "token is invalid or expired",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	c.Locals("id", claims.Id)
	c.Locals("email", claims.Issuer)
	return c.Next()
}
