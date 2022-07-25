package handler

import (
	"github.com/pestanko/gothy-mini/pkg/rest/rest_utils"
	"net/http"
)

// HandleHealth return OK status response if the application is ready
func HandleHealth() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rest_utils.WriteJSONResponse(w, http.StatusOK, statusDto{
			Status: "success",
		})
	}
}

type statusDto struct {
	Status string `json:"status"`
}
