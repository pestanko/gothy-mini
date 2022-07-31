package handler

import (
	"github.com/pestanko/gothy-mini/pkg/rest/resp"
	"net/http"
)

// HandleOAuth2Authorize handle OAuth 2.0 authorize request (Auth. code grant)
func HandleOAuth2Authorize() func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		return resp.MkErrResp(http.StatusNotImplemented, "TODO")
	}
}
