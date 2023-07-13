package iecho

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type context struct {
	echo.Context
}

func New(ctx echo.Context) *context {
	ctx.Echo().Validator = &customValidator{validator: validator.New()}
	return &context{Context: ctx}
}

func (c *context) ExtractResponse(err interface{}) *response {
	return err.(*response)
}

func (c *context) Success(data interface{}) error {
	sr := success
	sr.Data = data

	return c.JSON(http.StatusOK, sr)
}

func (c *context) FailBind(err error) error {
	data := make(map[string]string, 0)

	var he *echo.HTTPError
	if e, ok := err.(*echo.HTTPError); ok {
		he = e
	}

	msgLine := strings.ReplaceAll(strings.Split(fmt.Sprintf("%v", he.Message), ":")[1], " ", "")
	msgVal := strings.Split(msgLine, ",")
	for _, rv := range msgVal {
		v := strings.Split(rv, "=")
		data[v[0]] = v[1]
	}

	delete(data, "offset")

	response := ErrorInvalidFieldType
	response.Data = data

	return c.JSON(response.HC, response)
}

func (c *context) FailValidate(err error) error {
	vemap := getValidationErrors(err)
	response := ErrorBadRequest
	response.Data = vemap

	return c.JSON(response.HC, response)
}

func (c *context) Fail(err error) error {
	response := c.ExtractResponse(err)
	return c.JSON(response.HC, response)
}

func (c *context) Log(err error) {
	c.Context.Logger().Errorf("error %s", err.Error())
}
