package oauth2

import (
	"fmt"
	"github.com/pestanko/gothy-mini/pkg/client"
)

func ClientCredValidator(
	clientGetter client.Getter,
) func(cred *ClientCredentials) (found *client.Client, err error) {
	return func(cred *ClientCredentials) (found *client.Client, err error) {
		found, err = clientGetter.GetSingleClient(client.Query{ClientID: cred.ID})
		if err != nil {
			return
		}
		if found.Type == client.TypePublic {
			// do not check the secret
			return
		}
		if cred.Secret == "" {
			return found, fmt.Errorf("no client secret provided for: %s", found.ClientId)
		}
		if cred.Secret != found.AuthConfig.ClientSecret {
			return found, fmt.Errorf("client secret is invalid for %s", found.ClientId)
		}
		return
	}
}
