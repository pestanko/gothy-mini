package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/pestanko/gothy-mini/pkg/client"
	"github.com/pestanko/gothy-mini/pkg/rest/resp"
	"github.com/pestanko/gothy-mini/pkg/rest/restutl"
	"net/http"
)

// HandleClientList get list of all available client
func HandleClientList(clientGetter client.Getter) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		result, err := clientGetter.GetAllClients()
		if err != nil {
			return resp.MkServerError(err)
		}
		restutl.WriteJSONResp(w, http.StatusOK, result)
		return nil
	}
}

// HandleClientGet get a single client
func HandleClientGet(clientGetter client.Getter) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		clientId := chi.URLParam(r, "clientId")
		result, err := clientGetter.GetSingleClient(client.Query{ClientId: clientId})
		if err != nil {
			return resp.MkServerError(err)
		}
		restutl.WriteJSONResp(w, http.StatusOK, result)
		return nil
	}
}
