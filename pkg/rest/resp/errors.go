package resp

import (
	"fmt"
	"net/http"
)

// ErrorResponse interface
type ErrorResponse interface {
	Status() int
	ErrorDto() interface{}
	Error() string
}

// ErrorResp representation
type ErrorResp struct {
	StatusCode int
	ErrDto     ErrorDto
}

func (e *ErrorResp) Status() int {
	return e.StatusCode
}

func (e *ErrorResp) ErrorDto() interface{} {
	return e.ErrDto
}

func (e *ErrorResp) Error() string {
	return fmt.Sprintf("status=%d %v", e.Status(), e.ErrorDto())
}

func MkErrResp(status int, description string) ErrorResponse {
	mapping := map[int]string{
		http.StatusBadRequest:           "invalid_request",
		http.StatusUnauthorized:         "unauthorized",
		http.StatusForbidden:            "forbidden",
		http.StatusUnsupportedMediaType: "invalid_request",
		http.StatusMethodNotAllowed:     "method_not_allowed",
	}

	errorString, ok := mapping[status]
	if !ok {
		errorString = "server_error"
	}

	return MkErrorResp(status, ErrorDto{
		Err:         errorString,
		Description: description,
	})
}

func MkErrorResp(status int, errDto ErrorDto) ErrorResponse {
	return &ErrorResp{
		StatusCode: status,
		ErrDto:     errDto,
	}
}

func MkRespErr(status int, err error) ErrorResponse {
	return MkErrResp(status, fmt.Sprintf("%v", err))
}

func MkServerError(err error) ErrorResponse {
	return MkRespErr(http.StatusInternalServerError, err)
}

// ErrorDto Err JSON representation
type ErrorDto struct {
	Err         string `json:"error"`
	Description string `json:"error_description"`
}

func (e *ErrorDto) Error() string {
	return fmt.Sprintf("error=\"%s\" error_description=\"%s\"", e.Err, e.Description)
}
