package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/the-NZA/DB_Lab1/backend/internal/config"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storer"
)

type MySQLStore struct {
	db *sql.DB
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

	db, err := sql.Open(c.DBType, c.DBURL)
	if err != nil {
		return nil, err
	}

	return &MySQLStore{
		db: db,
	}, nil
}
