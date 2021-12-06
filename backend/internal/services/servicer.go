package services

import "github.com/the-NZA/DB_Lab1/backend/internal/models"

type Servicer interface {
	BookService() BookServicer
	GenreService() GenreServicer
	AuthorService() AuthorServicer
	Close() error
}

type BookServicer interface {
	Get(string) (models.Book, error)
	Add(models.Book) (models.Book, error)
	Update(models.Book) (models.Book, error)
	Delete(string) error
	GetAll() ([]models.Book, error)
}

type AuthorServicer interface {
	Get(string) (models.Author, error)
	Add(models.Author) (models.Author, error)
	Update(models.Author) (models.Author, error)
	Delete(string) error
	GetAll() ([]models.Author, error)
}

type GenreServicer interface {
	Get(string) (models.Genre, error)
	Add(models.Genre) (models.Genre, error)
	Update(models.Genre) (models.Genre, error)
	Delete(string) error
	GetAll() ([]models.Genre, error)
}
