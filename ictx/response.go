package ictx

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

func (r *Response) Embed(data interface{}) *Response {
	r.Data = data
	return r
}

var (
	Success            = NewResponse(http.StatusOK, "00", "success")
	ErrorUnauthorized  = NewResponse(http.StatusUnauthorized, "10", "authorization problem")
	ErrorDBRecNotfound = NewResponse(http.StatusNoContent, "20", "record not found")
	ErrorDBDuplicate   = NewResponse(http.StatusConflict, "21", "error duplicate")
	ErrorDB            = NewResponse(http.StatusInternalServerError, "29", "database problem")
	ErrorUnknown       = NewResponse(http.StatusInternalServerError, "99", "unhandled error")
)

var (
	ErrorBadRequest       = NewResponse(http.StatusBadRequest, "400", "bad request")
	ErrorInvalidEmail     = NewResponse(http.StatusBadRequest, "400", "invalid email format") // move to usecase later
	ErrorRecordNotFound   = NewResponse(http.StatusNotFound, "404", "record not found")
	ErrorInternalServer   = NewResponse(http.StatusInternalServerError, "500", "internal server error")
	ErrorInvalidFieldType = NewResponse(http.StatusBadRequest, "400", "invalid field type")
	// ErrorDuplicate        = NewResponse(http.StatusFound, "302", "duplicate record")

)
