package handler

import (
	"github.com/pestanko/gothy-mini/pkg/rest/restutl"
	"net/http"
)

// HandleAuthLoginApiToken handle standard user api token login
func HandleAuthLoginApiToken() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		restutl.WriteJSONResp(w, http.StatusNotImplemented, statusDto{
			Status: "TODO",
		})
	}
}
