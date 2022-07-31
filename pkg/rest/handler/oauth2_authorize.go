package handler

import (
	"github.com/pestanko/gothy-mini/pkg/rest/restutl"
	"net/http"
)

// HandleOAuth2Authorize handle OAuth 2.0 authorize request (Auth. code grant)
func HandleOAuth2Authorize() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		restutl.WriteJSONResp(w, http.StatusNotImplemented, statusDto{
			Status: "TODO",
		})
	}
}
