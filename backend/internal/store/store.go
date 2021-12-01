package store

import (
	"errors"

	"github.com/the-NZA/DB_Lab1/backend/internal/config"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/mock"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/mysql"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storer"
)

var (
	ErrUnsupportedDBType = errors.New("Unsupported DB type")
	ErrStoreNilOrEmpty   = errors.New("Store must be initialized before usage")
)

func NewStore(config *config.Config) (storer.Storer, error) {
	switch config.DBType {
	case "mysql":
		return mysql.NewStore(config)
	case "mock":
		return mock.NewStore(config)
	default:
		return nil, ErrUnsupportedDBType
	}
}
