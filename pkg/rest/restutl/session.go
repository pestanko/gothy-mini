package restutl

import (
	"context"
	"fmt"
	"github.com/pestanko/gothy-mini/pkg/auth/session"
	"net/http"
	"time"
)

// SessionIDKey session ID
const SessionIDKey = "SESSION_ID"

// SessionKey for context
const SessionKey = "session"

// GetSessionIDFromRequest get session ID as a string
func GetSessionIDFromRequest(r *http.Request) (string, error) {
	sessCookie, err := r.Cookie(SessionIDKey)
	if err != nil {
		return "", fmt.Errorf("unable to extact the session cookie: %w", err)
	}

	return sessCookie.Value, nil
}

// GetSessionFromRequest get session from the request
func GetSessionFromRequest(r *http.Request, store session.Store) (sess *session.Session, err error) {
	sessId, err := GetSessionIDFromRequest(r)
	if err != nil {
		return
	}

	sess, err = store.FindById(sessId)

	return
}

// RequireSession require authentication using the session
func RequireSession(w http.ResponseWriter, r *http.Request, store session.Store) (sess *session.Session) {
	sess, err := GetSessionFromRequest(r, store)
	if err != nil {
		WriteErrResp(w, http.StatusUnauthorized, err)
		return nil
	}
	if sess == nil || !sess.IsValid(time.Now()) {
		WriteErrorResp(w, MkErrResp(http.StatusUnauthorized, "you are not authorized"))
		return
	}
	return
}

// GetSessionFromCtx
func GetSessionFromCtx(ctx context.Context) *session.Session {
	return ctx.Value(SessionKey).(*session.Session)
}
