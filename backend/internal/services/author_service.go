package services

import (
	"github.com/the-NZA/DB_Lab1/backend/internal/models"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storer"
)

type AuthorService struct {
	repository storer.AuthorRepository
}

func (b *AuthorService) Get(ID string) (models.Author, error) {
	return b.repository.Get(ID)
}

func (b *AuthorService) Add(book models.Author) (models.Author, error) {
	return models.Author{
		ID:      "2",
		Snippet: "This is new book",
		Deleted: false,
	}, nil
}

func (b *AuthorService) Update(book models.Author) (models.Author, error) {
	return models.Author{
		ID:      "1",
		Snippet: "This is first book updated",
		Deleted: false,
	}, nil
}

func (b *AuthorService) Delete(ID string) error {
	return nil
}

func (b *AuthorService) GetAll() ([]models.Author, error) {
	return []models.Author{
		{
			ID:      "1",
			Snippet: "This is first book",
			Deleted: false,
		},
		{
			ID:      "2",
			Snippet: "This is new book",
			Deleted: false,
		},
	}, nil
}
