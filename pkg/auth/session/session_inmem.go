package session

import "context"

type storeImpl struct {
	store map[string]Session
}

// NewStore get a new store implementation
func NewStore() Store {
	return &storeImpl{
		store: make(map[string]Session),
	}
}

func (s *storeImpl) Store(ctx context.Context, session Session) error {
	s.store[session.SessionID] = session
	return nil
}

func (s *storeImpl) Remove(ctx context.Context, session Session) error {
	delete(s.store, session.SessionID)
	return nil
}

func (s *storeImpl) FindById(sessionID string) (*Session, error) {
	sess, ok := s.store[sessionID]
	if !ok {
		return nil, nil
	}
	return &sess, nil
}
