package mock

import (
	"github.com/the-NZA/DB_Lab1/backend/internal/config"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storer"
)

type MockStore struct {
	books *BookReporsitory
}

func (s *MockStore) Books() storer.BookReporsitory {
	if s.books != nil {
		return s.books
	}

	s.books = &BookReporsitory{}

	return s.books
}

func (s *MockStore) Authors() storer.AuthorRepository {
	return nil
}

func (s *MockStore) Genres() storer.GenreRepository {
	return nil
}

func NewStore(c *config.Config) (storer.Storer, error) {
	if c == nil {
		return nil, config.ErrEmptyConfig
	}

	return &MockStore{}, nil
}
