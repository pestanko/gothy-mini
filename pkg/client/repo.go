package client

// Query to get a single client
type Query struct {
	ClientId string
}

// Getter interface with get method to get a user info
type Getter interface {
	GetSingleClient(query Query) (*Client, error)
	GetAllClients() ([]Client, error)
}
