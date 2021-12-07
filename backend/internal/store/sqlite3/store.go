package sqlite3

// "db_url": "file:sqlite_data/lab.db?foreign_keys=on",

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/the-NZA/DB_Lab1/backend/internal/config"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storer"
)

type SQLiteStore struct {
	db      *sqlx.DB
	books   storer.BookReporsitory
	authors storer.AuthorRepository
	genres  storer.GenreRepository
}

func (s *SQLiteStore) Books() storer.BookReporsitory {
	if s.books != nil {
		return s.books
	}

	s.books = &BookRepository{db: s.db}

	return s.books
}

func (s *SQLiteStore) Authors() storer.AuthorRepository {
	return nil
}

func (s *SQLiteStore) Genres() storer.GenreRepository {
	return nil
}

func (s *SQLiteStore) Close() error {
	if s == nil {
		return storer.ErrStoreNilOrEmpty
	}

	return s.db.Close()
}

func NewStore(c *config.Config) (storer.Storer, error) {
	if c == nil {
		return nil, config.ErrEmptyConfig
	}

	dburl := fmt.Sprintf("file:%s?foreign_keys=on", c.DBURL)

	db, err := sqlx.Connect(c.DBType, dburl)
	if err != nil {
		return nil, err
	}

	return &SQLiteStore{
		db: db,
	}, nil
}
