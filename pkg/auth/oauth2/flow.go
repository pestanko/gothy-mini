package oauth2

import (
	"context"
	"fmt"
	"github.com/pestanko/gothy-mini/pkg/client"
	"time"
)

// FlowParams params for the flow
type FlowParams struct {
	ClientCredentials ClientCredentials `json:"client_credentials"`
	Client            *client.Client    `json:"client"`
	Scopes            []string          `json:"scopes"`
	Additional        map[string]string `json:"additional"`
}

type Flow interface {
	Process(ctx context.Context, params *FlowParams) (Tokens, error)
}

type ClientCredentials struct {
	ID     string `json:"client_id"`
	Secret string `json:"client_secret"`
}

type Tokens struct {
	Access    string `json:"access_token"`
	Refresh   string `json:"refresh_token"`
	Id        string `json:"id_token"`
	ExpiresIn string `json:"expires_in"`
	TokenType string `json:"token_type"`
}

func MkTokens(access, refresh, id string, exp time.Duration) Tokens {
	return Tokens{
		Access:    access,
		Refresh:   refresh,
		Id:        id,
		ExpiresIn: fmt.Sprint(exp.Seconds()),
		TokenType: "Bearer",
	}
}
