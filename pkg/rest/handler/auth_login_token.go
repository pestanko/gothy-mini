package handler

import (
	"github.com/pestanko/gothy-mini/pkg/rest/rest_utils"
	"net/http"
)

// HandleAuthLoginApiToken handle standard user api token login
func HandleAuthLoginApiToken() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rest_utils.WriteJSONResponse(w, http.StatusNotImplemented, statusDto{
			Status: "TODO",
		})
	}
}
