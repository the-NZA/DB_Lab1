package services

import (
	"github.com/the-NZA/DB_Lab1/backend/internal/models"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storer"
)

type AuthorService struct {
	repository storer.AuthorRepository
}

func (a *AuthorService) Get(ID string) (models.Author, error) {
	return a.repository.Get(ID)
}

func (a *AuthorService) Add(book models.Author) (models.Author, error) {
	// validation here
	return book, nil
}

func (a *AuthorService) Update(book models.Author) (models.Author, error) {
	// validation here

	return models.Author{
		ID:      "1",
		Snippet: "This is first book updated",
		Deleted: false,
	}, nil
}

func (a *AuthorService) Delete(ID string) error {
	return nil
}

func (a *AuthorService) GetAll() ([]models.Author, error) {
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
