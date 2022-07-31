package restutl

import (
	"net/http"
)

// ErrorResponse representation
type ErrorResponse struct {
	Status int
	Error  ErrorDto
}

func MkErrResp(status int, description string) *ErrorResponse {
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

	return &ErrorResponse{
		Status: status,
		Error: ErrorDto{
			Error:            errorString,
			ErrorDescription: description,
		},
	}
}

func MakeInvalidRequest(desc string) *ErrorResponse {
	return MkErrResp(http.StatusBadRequest, desc)
}
