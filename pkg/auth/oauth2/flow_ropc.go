package oauth2

import (
	"context"
	"github.com/pestanko/gothy-mini/pkg/auth/login"
	"github.com/pestanko/gothy-mini/pkg/user"
	"github.com/rs/zerolog/log"
)

type ROPCFlow struct {
	pwdLogin func(credentials login.PasswordCredentials) (*user.User, error)
}

func NewROPCFlow(
	pwdLogin func(credentials login.PasswordCredentials) (*user.User, error),
) Flow {
	return &ROPCFlow{
		pwdLogin: pwdLogin,
	}
}

func (f *ROPCFlow) Process(ctx context.Context, params *FlowParams) (Tokens, error) {
	userCred := f.extractUserCred(params)

	foundUser, err := f.pwdLogin(userCred)
	if err != nil {
		return Tokens{}, err
	}
	log.Info().Str("user", foundUser.Username).Msg("found user")

	return Tokens{}, nil
}

func (f *ROPCFlow) extractUserCred(params *FlowParams) login.PasswordCredentials {
	username := params.Additional["username"]
	password := params.Additional["password"]

	return login.PasswordCredentials{
		Username: username,
		Password: password,
	}
}
