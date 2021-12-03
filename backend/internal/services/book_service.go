package services

import (
	"time"

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
	return models.Book{
		ID:         "2",
		Title:      "New book",
		Snippet:    "This is new book",
		GenreID:    "2",
		PagesCnt:   2,
		Year:       time.Now(),
		BookLangID: "1",
		Deleted:    false,
	}, nil
}

func (b *BookService) Update(book models.Book) (models.Book, error) {
	return models.Book{
		ID:         "1",
		Title:      "First book updated",
		Snippet:    "This is first book updated",
		GenreID:    "1",
		PagesCnt:   1,
		Year:       time.Now(),
		BookLangID: "1",
		Deleted:    false,
	}, nil
}

func (b *BookService) Delete(ID string) error {
	return nil
}

func (b *BookService) GetAll() ([]models.Book, error) {
	return []models.Book{
		{
			ID:         "1",
			Title:      "First book",
			Snippet:    "This is first book",
			GenreID:    "1",
			PagesCnt:   1,
			Year:       time.Now(),
			BookLangID: "1",
			Deleted:    false,
		},
		{
			ID:         "2",
			Title:      "New book",
			Snippet:    "This is new book",
			GenreID:    "2",
			PagesCnt:   2,
			Year:       time.Now(),
			BookLangID: "1",
			Deleted:    false,
		},
	}, nil
}
