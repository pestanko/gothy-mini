package handler

import (
	"github.com/pestanko/gothy-mini/pkg/rest/rest_utils"
	"net/http"
)

// HandleAuthLoginCredentials handle standard credentials login
func HandleAuthLoginCredentials() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rest_utils.WriteJSONResponse(w, http.StatusNotImplemented, statusDto{
			Status: "TODO",
		})
	}
}
