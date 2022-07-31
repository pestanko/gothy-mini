package handler

import (
	"context"
	"github.com/pestanko/gothy-mini/pkg/auth/login"
	"github.com/pestanko/gothy-mini/pkg/auth/session"
	"github.com/pestanko/gothy-mini/pkg/rest/resp"
	"github.com/pestanko/gothy-mini/pkg/rest/restutl"
	"github.com/pestanko/gothy-mini/pkg/security"
	"github.com/pestanko/gothy-mini/pkg/user"
	"net/http"
	"time"
)

// HandleAuthLoginCredentials handle standard credentials login
func HandleAuthLoginCredentials(
	getter user.Getter,
	pwdHasher security.PasswordHasher,
	sessionStore session.Store,
) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {

		var credentialsDto struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if e := restutl.ReadJSONBody(r, &credentialsDto); e != nil {
			return e
		}

		doLogin := login.DoPasswordLogin(getter, pwdHasher)
		foundUser, err := doLogin(login.PasswordCredentials{
			Username: credentialsDto.Username,
			Password: credentialsDto.Password,
		})

		if err != nil {
			return resp.MkErrResp(http.StatusUnauthorized, "Invalid login")
		}

		sess := session.MakeSession(foundUser)
		if err := sessionStore.Store(context.Background(), sess); err != nil {
			return err
		}

		sendSessionCookie(w, sess)

		restutl.WriteJSONResp(w, http.StatusOK, foundUser)
		return nil
	}
}

func sendSessionCookie(w http.ResponseWriter, sess session.Session) {
	sessCookie := http.Cookie{
		Name:     restutl.SessionIDKey,
		Value:    sess.SessionID,
		Path:     "/",
		Expires:  time.Now().Add(8 * time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, &sessCookie)
}
