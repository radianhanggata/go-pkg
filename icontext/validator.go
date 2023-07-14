package icontext

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator"
	"github.com/iancoleman/strcase"

	"github.com/radianhanggata/go-pkg/iconst"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.Validator.Struct(i)
	if err != nil {
		return err
	}

	return nil
}

func getValidationErrors(err error) (sr iconst.Response) {
	var ve validator.ValidationErrors
	msg := make([]string, 0, len(ve))
	if errors.As(err, &ve) {
		for _, fe := range ve {
			msg = append(msg, strcase.ToSnake(fe.Field())+":"+describeTagError(fe))
		}
	}

	sr = iconst.Response{
		SC:      iconst.ErrorRequestValidationCode,
		Message: strings.Join(msg, ","),
	}

	return
}

func describeTagError(fe validator.FieldError) string {
	switch fe.Tag() {
	case "gte":
		return fmt.Sprintf("length must be equal or greater than %v", fe.Param())
	case "email":
		return "invalid email address"
	default:
		return fe.Tag()
	}
}
