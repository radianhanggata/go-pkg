package iconst

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const ErrorRequestValidationCode = "01"
const ErrorRequestBindCode = "02"

type Response struct {
	HttpStatus int         `json:"-"`
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func NewResponse(httpStatus int, code, message string) *Response {
	return &Response{
		HttpStatus: httpStatus,
		Code:       code,
		Message:    message,
		Data:       nil,
	}
}

func (r *Response) Error() string {
	return "unhandled error"
}

func (r *Response) ToHttpError() *echo.HTTPError {
	return &echo.HTTPError{
		Code:    r.HttpStatus,
		Message: r,
	}
}

var (
	Success               = NewResponse(http.StatusOK, "00", "success")
	ErrorBadRequest       = NewResponse(http.StatusBadRequest, "400", "bad request")
	ErrorInvalidEmail     = NewResponse(http.StatusBadRequest, "400", "invalid email format")
	ErrorRecordNotFound   = NewResponse(http.StatusNotFound, "404", "record not found")
	ErrorDuplicate        = NewResponse(http.StatusFound, "302", "duplicate record")
	ErrorInternalServer   = NewResponse(http.StatusInternalServerError, "500", "internal server error")
	ErrorInvalidFieldType = NewResponse(http.StatusBadRequest, "400", "invalid field type")
	ErrorUnauthorized     = NewResponse(http.StatusUnauthorized, "401", "unauthorized")
)
