package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/pestanko/gothy-mini/pkg/rest/rest_utils"
	"github.com/pestanko/gothy-mini/pkg/user"
	"net/http"
)

// HandleUserList get list of all available users
func HandleUserList(userGetter user.Getter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := userGetter.GetAllUsers()
		if err != nil {
			rest_utils.WriteErrResponse(w, http.StatusInternalServerError, err)
		} else {
			rest_utils.WriteJSONResp(w, http.StatusOK, users)
		}
	}
}

// HandleUserGet get a single user
func HandleUserGet(userGetter user.Getter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")
		result, err := userGetter.GetSingleUser(user.Query{Username: username})
		if err != nil {
			rest_utils.WriteErrResponse(w, http.StatusInternalServerError, err)
		} else {
			rest_utils.WriteJSONResp(w, http.StatusOK, result)
		}
	}
}
