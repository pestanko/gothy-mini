package login

import (
	"fmt"
	"github.com/pestanko/gothy-mini/pkg/security"
	"github.com/pestanko/gothy-mini/pkg/user"
)

var (
	InvalidPasswordError = fmt.Errorf("invalid password")
)

const randomPasswordHash = "$2a$14$ARw0H6VNqMQ5Whp8HaoefuWE.xjCdQ1rfVszILwJ3hgXRqF7L3ZVe"
const randomPassword = "SoMERandomPasswordToBeHasned_123456"

type PasswordCredentials struct {
	Username string
	Password string
}

func DoPasswordLogin(
	userGetter user.Getter,
	pwdHasher security.PasswordHasher,
) func(credentials PasswordCredentials) (*user.User, error) {
	return func(cred PasswordCredentials) (found *user.User, err error) {
		found, err = userGetter.GetSingleUser(user.Query{Username: cred.Username})
		if err != nil {
			pwdHasher.CheckPasswordHash(randomPassword, randomPasswordHash)
			return nil, err
		}

		if !pwdHasher.CheckPasswordHash(cred.Password, found.Cred.Password) {
			return nil, InvalidPasswordError
		}

		return
	}
}
