package handler

import (
	"github.com/pestanko/gothy-mini/pkg/rest/resp"
	"net/http"
)

// HandleAuthLoginApiToken handle standard user api token login
func HandleAuthLoginApiToken() func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		return resp.MkErrResp(http.StatusNotImplemented, "TODO")
	}
}
