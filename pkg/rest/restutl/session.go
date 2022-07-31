package restutl

import (
	"context"
	"fmt"
	"github.com/pestanko/gothy-mini/pkg/auth/session"
	"net/http"
)

// SessionIDKey session ID
const SessionIDKey = "SESSION_ID"

// SessionKey for context
const SessionKey = "session"

// ExtractSessionFromReqCookie get session from the request
func ExtractSessionFromReqCookie(r *http.Request, store session.Store) (sess *session.Session, err error) {
	sessId, err := getSessionIDFromRequest(r)
	if err != nil {
		return
	}

	sess, err = store.FindById(sessId)

	return
}

// GetSessionFromCtx session from the request context
func GetSessionFromCtx(ctx context.Context) *session.Session {
	value := ctx.Value(SessionKey)
	if value != nil {
		return value.(*session.Session)
	}
	return nil
}

// GetSessionFromReq get a session that is stored in the request context
func GetSessionFromReq(r *http.Request) *session.Session {
	return GetSessionFromCtx(r.Context())
}

// getSessionIDFromRequest get session ID as a string
func getSessionIDFromRequest(r *http.Request) (string, error) {
	sessCookie, err := r.Cookie(SessionIDKey)
	if err != nil {
		return "", fmt.Errorf("unable to extact the session cookie: %w", err)
	}

	return sessCookie.Value, nil
}
