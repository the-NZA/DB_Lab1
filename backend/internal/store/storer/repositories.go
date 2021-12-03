package storer

import "github.com/the-NZA/DB_Lab1/backend/internal/models"

type BookReporsitory interface {
	Get(ID string) (models.Book, error)
}

type AuthorRepository interface {
	Get(ID string) (models.Author, error)
}

type GenreRepository interface {
	Get(ID string) (models.Genre, error)
}
