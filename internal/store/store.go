package store

import "errors"

type IStore interface {
	Get(key string) (string, error)

	Set(key, value string) error
}

type memoryStore struct {
	data map[string]string
}

func NewMemoryStore() IStore {
	return &memoryStore{
		data: make(map[string]string),
	}
}

func (s *memoryStore) Get(key string) (string, error) {
	return s.data[key], nil
}
func (s *memoryStore) Set(key, value string) error {
	if _, ok := s.data[key]; ok {
		return errors.New("key already exists")
	}
	s.data[key] = value
	return nil
}
