package session

import (
	"context"
	"github.com/pestanko/gothy-mini/pkg/security"
	"github.com/pestanko/gothy-mini/pkg/user"
	"time"
)

const sessionExpiration = 8 * time.Hour

// Session representation
type Session struct {
	SessionID string    `json:"session_id"`
	Username  string    `json:"username"`
	UserType  user.Type `json:"user_type"`
	ExpireAt  time.Time `json:"expire_at"`
}

func (s *Session) IsValid(at time.Time) bool {
	return s.ExpireAt.After(at)
}

func MakeSession(u *user.User) Session {
	return Session{
		SessionID: security.RandomString(32),
		Username:  u.Username,
		UserType:  u.Type,
		ExpireAt:  time.Now().Add(sessionExpiration),
	}
}

// Store representation of the session store
type Store interface {
	Store(ctx context.Context, session Session) error
	Remove(ctx context.Context, session Session) error
	FindById(sessionID string) (*Session, error)
}
