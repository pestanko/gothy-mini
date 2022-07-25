package user

// Type is a string enum
type Type string

const (
	// TypeAdmin is an admin user type
	TypeAdmin Type = "admin"
	// TypeSystem is a system user type
	TypeSystem Type = "system"
	// TypeUser is a normal user type
	TypeUser Type = "user"
)

// User entity
type User struct {
	Username string      `yaml:"username" json:"username"`
	Email    string      `yaml:"email" json:"email"`
	Name     string      `yaml:"name" json:"name"`
	Type     Type        `yaml:"type" json:"type"`
	Cred     Credentials `yaml:"credentials" json:"credentials"`
}

// Credentials represents a user credentials
type Credentials struct {
	Password string  `yaml:"password" json:"-"`
	Tokens   []Token `yaml:"tokens" json:"tokens"`
}

// Token represents an access token for the user
type Token struct {
	Name   string   `yaml:"name" json:"name"`
	Value  string   `yaml:"value" json:"-"`
	Scopes []string `yaml:"scopes" json:"scopes"`
}
