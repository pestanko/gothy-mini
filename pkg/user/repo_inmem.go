package user

import (
	"fmt"
)

// NewGetter for getting users
func NewGetter(users []User) Getter {
	return &repoInMemory{
		users: users,
	}
}

type repoInMemory struct {
	users []User
}

func (r *repoInMemory) GetSingleUser(query Query) (*User, error) {
	for _, user := range r.getAll() {
		if query.Username != "" && user.Username == query.Username {
			return &user, nil
		}
		if query.Email != "" && user.Email == query.Email {
			return &user, nil
		}
	}

	return nil, fmt.Errorf("unable to get any user for: %v", query)
}

func (r *repoInMemory) GetAllUsers() ([]User, error) {
	return r.getAll(), nil
}

func (r *repoInMemory) getAll() []User {
	return r.users
}
