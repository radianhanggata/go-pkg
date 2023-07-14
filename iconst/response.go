package iconst

import "net/http"

const ErrorRequestValidationCode = "01"
const ErrorRequestBindCode = "02"

type Response struct {
	HC      int         `json:"-"`
	SC      string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(hc int, code, message string) *Response {
	return &Response{
		HC:      hc,
		SC:      code,
		Message: message,
		Data:    nil,
	}
}

func (r *Response) Error() string {
	return r.Message
}

var (
	Success               = NewResponse(http.StatusOK, "00", "success")
	ErrorBadRequest       = NewResponse(http.StatusBadRequest, "400", "bad request")
	ErrorInvalidEmail     = NewResponse(http.StatusBadRequest, "400", "invalid email format")
	ErrorRecordNotFound   = NewResponse(http.StatusNotFound, "404", "record not found")
	ErrorDuplicate        = NewResponse(http.StatusFound, "302", "duplicate record")
	ErrorInternalServer   = NewResponse(http.StatusInternalServerError, "500", "internal server error")
	ErrorInvalidFieldType = NewResponse(http.StatusBadRequest, "400", "invalid field type")
)
