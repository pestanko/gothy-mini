package rest_utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"strings"
)

// ErrorDto Error JSON representation
type ErrorDto struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// WriteErrorResp helper
func WriteErrorResp(w http.ResponseWriter, resp *ErrorResponse) {
	log.Warn().
		Str("error", resp.Error.Error).
		Str("error_desc", resp.Error.ErrorDescription).
		Int("status_code", resp.Status).
		Msg("Returning the error response")

	WriteJSONResp(w, resp.Status, resp.Error)
}

// WriteErrResponse helper
func WriteErrResponse(w http.ResponseWriter, code int, err error) {
	WriteErrorResp(w, MkErrResp(code, fmt.Sprintf("%v", err)))
}

// WriteJSONResp helper
func WriteJSONResp(w http.ResponseWriter, code int, resp interface{}) {
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

func ReadJSONBody(r *http.Request, out interface{}) *ErrorResponse {
	// Inspiration: https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body

	if response := checkContentType(r, "application/json"); response != nil {
		return response
	}

	// Setup the decoder and call the DisallowUnknownFields() method on it.
	// This will cause Decode() to return a "json: unknown field ..." error
	// if it encounters any extra unexpected fields in the JSON. Strictly
	// speaking, it returns an error for "keys which do not match any
	// non-ignored, exported fields in the destination".
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(out); err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		// Catch any syntax errors in the JSON and send an error message
		// which interpolates the location of the problem to make it
		// easier for the client to fix.
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return MakeInvalidRequest(msg)

		// In some circumstances Decode() may also return an
		// io.ErrUnexpectedEOF error for syntax errors in the JSON. There
		// is an open issue regarding this at
		// https://github.com/golang/go/issues/25956.
		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			return MakeInvalidRequest(msg)

		// Catch any type errors, like trying to assign a string in the
		// JSON request body to a int field in our Person struct. We can
		// interpolate the relevant field name and position into the error
		// message to make it easier for the client to fix.
		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return MakeInvalidRequest(msg)

		// Catch the error caused by extra unexpected fields in the request
		// body. We extract the field name from the error message and
		// interpolate it in our custom error message. There is an open
		// issue at https://github.com/golang/go/issues/29035 regarding
		// turning this into a sentinel error.
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return MakeInvalidRequest(msg)

		// An io.EOF error is returned by Decode() if the request body is
		// empty.
		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return MakeInvalidRequest(msg)

		// Catch the error caused by the request body being too large. Again
		// there is an open issue regarding turning this into a sentinel
		// error at https://github.com/golang/go/issues/30715.
		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return MkErrResp(http.StatusRequestEntityTooLarge, msg)

		// Otherwise default to logging the error and sending a 500 Internal
		// Server Error response.
		default:
			return MkErrResp(http.StatusInternalServerError, "Oops something went wrong")
		}
	}

	return nil
}

func checkContentType(r *http.Request, expected string) *ErrorResponse {
	headerValue := r.Header.Get("Content-Type")
	if headerValue != "" && headerValue != expected {
		msg := fmt.Sprintf("Content-Type header is not \"%s\"", expected)
		return MkErrResp(http.StatusUnsupportedMediaType, msg)

	}
	return nil
}
