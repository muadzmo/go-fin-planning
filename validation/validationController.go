package validation

import (
	"github.com/go-playground/validator"
)

func Validating(data interface{}) error {
	validate := validator.New()
	err := validate.Struct(data)
	return err
}
