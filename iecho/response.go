package iecho

import "net/http"

type response struct {
	HC      int         `json:"-"`
	SC      string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(hc int, code, message string) *response {
	return &response{
		HC:      hc,
		SC:      code,
		Message: message,
		Data:    nil,
	}
}

func (r *response) Error() string {
	return r.Message
}

var success = NewResponse(http.StatusOK, "00", "success")

var (
	ErrorBadRequest        = NewResponse(http.StatusBadRequest, "400", "bad request")
	ErrorInvalidEmail      = NewResponse(http.StatusBadRequest, "400", "invalid email format")
	ErrorInvalidCredential = NewResponse(http.StatusUnauthorized, "401", "invalid username or secret")
	ErrorRecordNotFound    = NewResponse(http.StatusNotFound, "404", "record not found")
	ErrorDuplicate         = NewResponse(http.StatusFound, "302", "duplicate record")
	ErrorInternalServer    = NewResponse(http.StatusInternalServerError, "500", "internal server error")
	ErrorInvalidFieldType  = NewResponse(http.StatusBadRequest, "400", "invalid field type")
)
