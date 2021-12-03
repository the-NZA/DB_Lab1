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

func (a *AuthorService) Add(author models.Author) (models.Author, error) {
	// validation here
	return a.repository.Add(author)
}

func (a *AuthorService) Update(author models.Author) (models.Author, error) {
	// validation here
	return a.repository.Update(author)
}

func (a *AuthorService) Delete(ID string) error {
	//validation
	return a.repository.Delete(ID)
}

func (a *AuthorService) GetAll() ([]models.Author, error) {
	return a.repository.GetAll()
}
