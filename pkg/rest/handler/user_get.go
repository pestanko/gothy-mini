package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/pestanko/gothy-mini/pkg/rest/resp"
	"github.com/pestanko/gothy-mini/pkg/rest/restutl"
	"github.com/pestanko/gothy-mini/pkg/user"
	"net/http"
)

// HandleUserList get list of all available users
func HandleUserList(userGetter user.Getter) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		users, err := userGetter.GetAllUsers()
		if err != nil {
			return resp.MkServerError(err)
		}
		restutl.WriteJSONResp(w, http.StatusOK, users)
		return nil
	}
}

// HandleUserGet get a single user
func HandleUserGet(userGetter user.Getter) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		username := chi.URLParam(r, "username")
		result, err := userGetter.GetSingleUser(user.Query{Username: username})
		if err != nil {
			return resp.MkServerError(err)
		}

		restutl.WriteJSONResp(w, http.StatusOK, result)
		return nil
	}
}
