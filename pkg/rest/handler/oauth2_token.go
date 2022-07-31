package handler

import (
	"github.com/pestanko/gothy-mini/pkg/rest/restutl"
	"github.com/pestanko/gothy-mini/pkg/user"
	"net/http"
)

// HandleOAuth2Token handle OAuth 2.0 authorize request (Auth. code grant)
func HandleOAuth2Token(userGetter user.Getter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			restutl.WriteErrResp(w, http.StatusBadRequest, err)
			return
		}

		grantType := r.PostForm.Get("grant_type")
		if grantType == "" {
			restutl.RequiredParamMissing(w, "grant_type")
			return
		}

		switch grantType {
		case "password": // resource owner password credentials

		default:
			restutl.InvalidParamM(w, "grant_type", grantType)
			return
		}

		restutl.WriteJSONResp(w, http.StatusNotImplemented, statusDto{
			Status: "TODO",
		})
	}
}
