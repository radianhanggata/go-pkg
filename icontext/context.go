package icontext

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	"github.com/radianhanggata/go-pkg/iconst"
)

type context struct {
	echo.Context
}

func New(ctx echo.Context) *context {
	return &context{Context: ctx}
}

func (c *context) Success(data interface{}) error {
	sr := iconst.Success
	sr.Data = data

	return c.JSON(http.StatusOK, sr)
}

func (c *context) Fail(err error) error {
	switch err.(type) {
	case *echo.HTTPError:
		return c.failBind(err)
	case validator.ValidationErrors:
		return c.failValidate(err)
	default:
		response := extractResponse(err)
		return c.JSON(response.HC, response)
	}
}

func (c *context) Log(err error) {
	c.Context.Logger().Errorf("error %s", err.Error())
}

func extractResponse(err interface{}) *iconst.Response {
	return err.(*iconst.Response)
}

func (c *context) failBind(err error) error {
	response := iconst.ErrorInvalidFieldType

	he := err.(*echo.HTTPError)

	data := iconst.Response{
		SC:      iconst.ErrorRequestBindCode,
		Message: fmt.Sprintf("%v", he.Message),
	}

	response.Data = data

	return c.JSON(response.HC, response)
}

func (c *context) failValidate(err error) error {
	vemap := getValidationErrors(err)
	response := iconst.ErrorBadRequest
	response.Data = vemap

	return c.JSON(response.HC, response)
}
