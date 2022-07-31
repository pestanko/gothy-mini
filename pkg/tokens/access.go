package tokens

import (
	"github.com/pestanko/gothy-mini/pkg/client"
	"github.com/pestanko/gothy-mini/pkg/user"
)

type AccessTokens struct {
}

type TokenParams struct {
	User   *user.User
	Client *client.Client
	Scopes []string
	Flow   string
}

func (t *AccessTokens) MakeToken(params TokenParams) (string, error) {

}
