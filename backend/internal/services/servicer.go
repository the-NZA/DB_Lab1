package services

import "github.com/the-NZA/DB_Lab1/backend/internal/models"

type Servicer interface {
	BookService() BookServicer
	// AuthorService() AuthorServicer
	// GenreService() GenreServicer
}

type BookServicer interface {
	Get() models.Book
	Add() (models.Book, error)
	Update() (models.Book, error)
	Delete() error
	FindAll() ([]models.Book, error)
}

type AuthorServicer interface {
	Get()
	Add()
	Update()
	Delete()
	FindAll()
}

type GenreServicer interface {
	Get()
	Add()
	Update()
	Delete()
	FindAll()
}
