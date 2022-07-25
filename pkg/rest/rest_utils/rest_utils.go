package rest_utils

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
)

type ErrorDto struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// WriteErrorResponse helper
func WriteErrorResponse(w http.ResponseWriter, code int, err ErrorDto) {
	log.Warn().
		Str("error", err.Error).
		Str("error_desc", err.ErrorDescription).
		Int("code", code).
		Msg("Returning the error response")

	WriteJSONResponse(w, code, err)
}

// WriteErrResponse helper
func WriteErrResponse(w http.ResponseWriter, code int, err error) {
	WriteErrorResponse(w, code, ErrorDto{
		Error:            "server_error",
		ErrorDescription: fmt.Sprintf("%v", err),
	})
}

// WriteJSONResponse helper
func WriteJSONResponse(w http.ResponseWriter, code int, resp interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Error().Err(err).Msg("Error happened in JSON marshal")
	}
	if _, err := w.Write(jsonResp); err != nil {
		log.Error().Err(err).Msg("Error writing response")
	}
}

// WriteUnsupportedHTTPMethod helper
func WriteUnsupportedHTTPMethod(w http.ResponseWriter, method string) {
	WriteErrorResponse(w, http.StatusMethodNotAllowed, ErrorDto{
		"unsupported_http_method",
		"Unsupported http method: " + method,
	})
}
