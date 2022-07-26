package handler

import (
	"github.com/pestanko/gothy-mini/pkg/auth/login"
	"github.com/pestanko/gothy-mini/pkg/rest/rest_utils"
	"github.com/pestanko/gothy-mini/pkg/security"
	"github.com/pestanko/gothy-mini/pkg/user"
	"net/http"
)

// HandleAuthLoginCredentials handle standard credentials login
func HandleAuthLoginCredentials(
	getter user.Getter,
	pwdHasher security.PasswordHasher,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var credentialsDto struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if e := rest_utils.ReadJSONBody(r, &credentialsDto); e != nil {
			rest_utils.WriteErrorResp(w, e)
			return
		}

		doLogin := login.DoPasswordLogin(getter, pwdHasher)
		foundUser, err := doLogin(login.PasswordCredentials{
			Username: credentialsDto.Username,
			Password: credentialsDto.Username,
		})

		if err != nil {
			rest_utils.WriteErrorResp(w, rest_utils.MkErrResp(http.StatusUnauthorized, "Invalid login"))
			return
		}

		rest_utils.WriteJSONResp(w, http.StatusOK, foundUser)
	}
}
