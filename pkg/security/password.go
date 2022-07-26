package security

import (
	"golang.org/x/crypto/bcrypt"
)

// PasswordHasher - shared interface for password hashing functionality
type PasswordHasher interface {
	// HashPassword Hash user's password
	HashPassword(password string) (string, error)
	// CheckPasswordHash check the hash against the password
	CheckPasswordHash(password, hash string) bool
}

func NewPasswordHasher() PasswordHasher {
	return &bcryptPwdHasher{
		cost: 14,
	}
}

type bcryptPwdHasher struct {
	cost int
}

// HashPassword Hash user's password
func (h *bcryptPwdHasher) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	return string(bytes), err
}

func (h *bcryptPwdHasher) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
