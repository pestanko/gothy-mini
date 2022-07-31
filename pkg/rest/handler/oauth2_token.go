package handler

import (
	"github.com/pestanko/gothy-mini/pkg/rest/resp"
	"github.com/pestanko/gothy-mini/pkg/user"
	"net/http"
)

// HandleOAuth2Token handle OAuth 2.0 authorize request (Auth. code grant)
func HandleOAuth2Token(userGetter user.Getter) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		if err := r.ParseForm(); err != nil {
			return resp.MkRespErr(http.StatusBadRequest, err)
		}

		grantType := r.PostForm.Get("grant_type")
		if grantType == "" {
			return resp.MkParamMissing("grant_type")
		}

		switch grantType {
		case "password": // resource owner password credentials

		default:
			return resp.MkParamInvalid("grant_type", grantType)
		}

		return resp.MkErrResp(http.StatusNotImplemented, "TODO")
	}
}
