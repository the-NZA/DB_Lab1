package services

import (
	"github.com/the-NZA/DB_Lab1/backend/internal/models"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storer"
)

type GenreService struct {
	repository storer.GenreRepository
}

func (b *GenreService) Get(ID string) (models.Genre, error) {
	return b.repository.Get(ID)
}

func (b *GenreService) Add(book models.Genre) (models.Genre, error) {
	return models.Genre{
		ID:      "2",
		Title:   "New book",
		Snippet: "This is new book",
		Deleted: false,
	}, nil
}

func (b *GenreService) Update(book models.Genre) (models.Genre, error) {
	return models.Genre{
		ID:      "1",
		Title:   "First book updated",
		Snippet: "This is first book updated",
		Deleted: false,
	}, nil
}

func (b *GenreService) Delete(ID string) error {
	return nil
}

func (b *GenreService) GetAll() ([]models.Genre, error) {
	return []models.Genre{
		{
			ID:      "1",
			Title:   "First book",
			Snippet: "This is first book",
			Deleted: false,
		},
		{
			ID:      "2",
			Title:   "New book",
			Snippet: "This is new book",
			Deleted: false,
		},
	}, nil
}
