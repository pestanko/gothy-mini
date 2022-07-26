package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/pestanko/gothy-mini/pkg/client"
	"github.com/pestanko/gothy-mini/pkg/rest/rest_utils"
	"net/http"
)

// HandleClientList get list of all available client
func HandleClientList(clientGetter client.Getter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := clientGetter.GetAllClients()
		if err != nil {
			rest_utils.WriteErrResponse(w, http.StatusInternalServerError, err)
		} else {
			rest_utils.WriteJSONResp(w, http.StatusOK, users)
		}
	}
}

// HandleClientGet get a single client
func HandleClientGet(clientGetter client.Getter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		clientId := chi.URLParam(r, "clientId")
		result, err := clientGetter.GetSingleClient(client.Query{ClientId: clientId})
		if err != nil {
			rest_utils.WriteErrResponse(w, http.StatusInternalServerError, err)
		} else {
			rest_utils.WriteJSONResp(w, http.StatusOK, result)
		}
	}
}
