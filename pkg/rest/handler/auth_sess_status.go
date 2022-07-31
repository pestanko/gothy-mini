package handler

import (
	"github.com/pestanko/gothy-mini/pkg/rest/resp"
	"github.com/pestanko/gothy-mini/pkg/rest/restutl"
	"net/http"
)

func HandleAuthSessionStatus() func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		sess := restutl.GetSessionFromReq(r)
		if sess == nil {
			return resp.MkErrResp(http.StatusUnauthorized, "no valid session found")
		}

		restutl.WriteJSONResp(w, http.StatusOK, sess)
		return nil
	}
}
