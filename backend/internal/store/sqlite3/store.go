package sqlite3

// "db_url": "file:sqlite_data/lab.db?foreign_keys=on",

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/the-NZA/DB_Lab1/backend/internal/config"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storer"
)

type SQLiteStore struct {
	db *sql.DB
}

func (s *SQLiteStore) Books() storer.BookReporsitory {
	return nil
}

func (s *SQLiteStore) Authors() storer.AuthorRepository {
	return nil
}

func (s *SQLiteStore) Genres() storer.GenreRepository {
	return nil
}

func NewStore(c *config.Config) (storer.Storer, error) {
	if c == nil {
		return nil, config.ErrEmptyConfig
	}

	dburl := fmt.Sprintf("file:%s?foreign_keys=on", c.DBURL)

	db, err := sql.Open(c.DBType, dburl)
	if err != nil {
		return nil, err
	}

	return &SQLiteStore{
		db: db,
	}, nil
}
