package handler

import (
	"github.com/pestanko/gothy-mini/pkg/rest/restutl"
	"net/http"
)

func HandleAuthSessionStatus() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sess := restutl.GetSessionFromReq(r)
		if sess == nil {
			restutl.WriteErrorResp(w, restutl.MkErrResp(http.StatusUnauthorized, "no valid session found"))
			return
		}

		restutl.WriteJSONResp(w, http.StatusOK, sess)
	}
}
