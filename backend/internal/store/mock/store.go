package mock

import (
	"github.com/the-NZA/DB_Lab1/backend/internal/config"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storer"
)

type MockStore struct{}

func (s *MockStore) Books() storer.BookReporsitory {
	return nil
}

func (s *MockStore) Authors() storer.AuthorRepository {
	return nil
}

func (s *MockStore) Genres() storer.GenreRepository {
	return nil
}

func NewStore(config *config.Config) (storer.Storer, error) {
	return &MockStore{}, nil
}
