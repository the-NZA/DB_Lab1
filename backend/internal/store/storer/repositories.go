package storer

import "github.com/the-NZA/DB_Lab1/backend/internal/models"

type BookReporsitory interface {
	Get(string) (models.Book, error)
	Add(models.Book) (models.Book, error)
	Update(models.Book) (models.Book, error)
	Delete(string) error
	GetAll() ([]models.Book, error)
}

type AuthorRepository interface {
	Get(string) (models.Author, error)
}

type GenreRepository interface {
	Get(string) (models.Genre, error)
}
