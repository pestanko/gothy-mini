package handler

import (
	"github.com/pestanko/gothy-mini/pkg/auth/session"
	"github.com/pestanko/gothy-mini/pkg/rest/restutl"
	"net/http"
)

func HandleAuthSessionStatus(
	sessStore session.Store,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sess := restutl.RequireSession(w, r, sessStore)
		if sess == nil {
			return
		}

		restutl.WriteJSONResp(w, http.StatusOK, sess)
	}
}
