package handler

import (
	"github.com/pestanko/gothy-mini/pkg/rest/rest_utils"
	"net/http"
)

// HandleOAuth2Token handle OAuth 2.0 authorize request (Auth. code grant)
func HandleOAuth2Token() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rest_utils.WriteJSONResp(w, http.StatusNotImplemented, statusDto{
			Status: "TODO",
		})
	}
}
