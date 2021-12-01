package services

import (
	"github.com/the-NZA/DB_Lab1/backend/internal/config"
	"github.com/the-NZA/DB_Lab1/backend/internal/store"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storer"
)

var ()

type Services struct {
	store  storer.Storer
	config *config.Config
	books  *BookService
}

func (s *Services) BookService() BookServicer {
	if s.books != nil {
		return s.books
	}

	s.books = &BookService{
		repository: s.store.Books(),
	}

	return s.books
}

func (s *Services) AuthorService() AuthorServicer {
	return nil
}

func (s *Services) GenreService() GenreServicer {
	return nil
}

// NewServices create initiale new services structure
func NewServices(c *config.Config, s storer.Storer) (Servicer, error) {
	if c == nil {
		return nil, config.ErrEmptyConfig
	}

	if s == nil {
		return nil, store.ErrStoreNilOrEmpty
	}

	return &Services{
		store:  s,
		config: c,
	}, nil
}
