package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/the-NZA/DB_Lab1/backend/internal/config"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storer"
)

type MySQLStore struct {
	db *sqlx.DB
}

func (s *MySQLStore) Books() storer.BookReporsitory {
	return nil
}

func (s *MySQLStore) Authors() storer.AuthorRepository {
	return nil
}

func (s *MySQLStore) Genres() storer.GenreRepository {
	return nil
}

func (s *MySQLStore) BooksAuthors() storer.BookAuthorRepository {
	return nil
}

func (s *MySQLStore) Close() error {
	if s == nil {
		return storer.ErrStoreNilOrEmpty
	}

	return s.db.Close()
}

func NewStore(c *config.Config) (storer.Storer, error) {
	if c == nil {
		return nil, config.ErrEmptyConfig
	}

	db, err := sqlx.Connect(c.DBType, c.DBURL)
	if err != nil {
		return nil, err
	}

	return &MySQLStore{
		db: db,
	}, nil
}
