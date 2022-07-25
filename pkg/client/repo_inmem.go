package client

import "fmt"

// NewGetter for getting clients
func NewGetter(clients []Client) Getter {
	return &repoInMemory{
		clients: clients,
	}
}

type repoInMemory struct {
	clients []Client
}

func (r *repoInMemory) GetSingleClient(query Query) (*Client, error) {
	for _, client := range r.getAll() {
		if query.ClientId != "" && client.ClientId == query.ClientId {
			return &client, nil
		}
	}

	return nil, fmt.Errorf("unable to get any user for: %v", query)
}

func (r *repoInMemory) GetAllClients() ([]Client, error) {
	return r.getAll(), nil
}

func (r *repoInMemory) getAll() []Client {
	return r.clients
}
