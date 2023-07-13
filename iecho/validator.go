package iecho

import (
	"errors"

	"github.com/go-playground/validator"
	"github.com/iancoleman/strcase"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)
	if err != nil {
		return err
	}

	return nil
}

func getValidationErrors(err error) map[string]string {
	out := make(map[string]string, 0)
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, fe := range ve {
			out[strcase.ToSnake(fe.Field())] = fe.Tag()
		}
	}
	return out
}
