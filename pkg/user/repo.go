package user

// Query to get a single user
type Query struct {
	Username string
	Email    string
}

// Getter interface with get method to get a user info
type Getter interface {
	GetSingleUser(query Query) (*User, error)
	GetAllUsers() ([]User, error)
}
