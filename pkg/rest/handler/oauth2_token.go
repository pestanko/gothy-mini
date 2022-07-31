package handler

import (
	"fmt"
	"github.com/pestanko/gothy-mini/pkg/auth/oauth2"
	"github.com/pestanko/gothy-mini/pkg/client"
	"github.com/pestanko/gothy-mini/pkg/rest/resp"
	"github.com/pestanko/gothy-mini/pkg/rest/restutl"
	"net/http"
	"strings"
)

// HandleOAuth2Token handle OAuth 2.0 authorize request (Auth. code grant)
func HandleOAuth2Token(
	clientCredValidator func(cred *oauth2.ClientCredentials) (*client.Client, error),
	flows oauth2.Flows,
) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		if err := r.ParseForm(); err != nil {
			return resp.MkRespErr(http.StatusBadRequest, err)
		}

		clientCred := parseClientCred(r)
		if clientCred.ID == "" {
			return resp.MkParamMissing("client_id")
		}

		foundClient, err := clientCredValidator(&clientCred)
		if err != nil {
			return resp.MkErrorResp(http.StatusUnauthorized, resp.ErrorDto{
				Err:         "unauthorized_client",
				Description: fmt.Sprintf("%v", err),
			})
		}

		flowParams := oauth2.FlowParams{
			ClientCredentials: clientCred,
			Client:            foundClient,
			Additional:        make(map[string]string),
		}

		flow, err := getFlow(flows, r, &flowParams)
		if err != nil {
			return err
		}

		tokens, err := flow.Process(r.Context(), &flowParams)
		if err != nil {
			return err
		}

		restutl.WriteJSONResp(w, http.StatusOK, tokens)

		return nil
	}
}

func getFlow(flows oauth2.Flows, r *http.Request, params *oauth2.FlowParams) (oauth2.Flow, error) {
	grantType := r.PostForm.Get("grant_type")
	if grantType == "" {
		return nil, resp.MkParamMissing("grant_type")
	}
	grantType = strings.ToLower(grantType)
	switch grantType {
	case "password":
		err := processRopcFlowParams(r, params)
		if err != nil {
			return nil, err
		}
	default:
		return nil, resp.MkParamInvalid("grant_type", grantType)
	}
	return flows.GetFlow(grantType), nil
}

func processRopcFlowParams(r *http.Request, params *oauth2.FlowParams) (err error) {
	params.Additional["username"], err = requireParam(r, "username")
	if err != nil {
		return err
	}
	params.Additional["password"], err = requireParam(r, "password")
	if err != nil {
		return err
	}

	params.Additional["scope"], err = requireParam(r, "scope")
	if err != nil {
		return err
	}

	return nil
}

func parseClientCred(r *http.Request) (cred oauth2.ClientCredentials) {
	cred = parseFromAuthHeader(r)

	if cred.ID == "" {
		if clientID := r.PostForm.Get("client_id"); clientID != "" {
			cred.ID = clientID
		}
	}

	if cred.Secret == "" {
		if secret := r.PostForm.Get("client_secret"); secret != "" {
			cred.Secret = secret
		}
	}

	return cred
}

func parseFromAuthHeader(r *http.Request) (cred oauth2.ClientCredentials) {
	username, password, ok := r.BasicAuth()
	if ok {
		cred.ID = username
		cred.Secret = password
	}

	return
}

func requireParam(r *http.Request, name string) (string, error) {
	value := r.PostForm.Get(name)
	if value == "" {
		return "", resp.MkParamMissing(name)
	}
	return value, nil
}
