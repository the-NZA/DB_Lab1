package services

import (
	"github.com/the-NZA/DB_Lab1/backend/internal/models"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storer"
)

type BookService struct {
	repository storer.BookReporsitory
}

func (b *BookService) Get(ID string) (models.Book, error) {
	return b.repository.Get(ID)
}

func (b *BookService) Add(book models.Book) (models.Book, error) {
	// validation here
	return b.repository.Add(book)
}

func (b *BookService) Update(book models.Book) (models.Book, error) {
	// validation here
	return b.repository.Update(book)
}

func (b *BookService) Delete(ID string) error {
	// validation here
	return b.repository.Delete(ID)
}

func (b *BookService) GetAll() ([]models.Book, error) {
	return b.repository.GetAll()
}
