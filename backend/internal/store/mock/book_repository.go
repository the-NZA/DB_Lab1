package mock

import (
	"fmt"
	"time"

	"github.com/the-NZA/DB_Lab1/backend/internal/models"
)

var books = []models.Book{
	{
		ID:         "1",
		Title:      "First",
		Snippet:    "First book description",
		GenreID:    "1",
		PagesCnt:   100,
		Year:       time.Now(),
		BookLangID: "1",
		Deleted:    false,
	},
	{
		ID:         "2",
		Title:      "Second",
		Snippet:    "Second book description",
		GenreID:    "2",
		PagesCnt:   120,
		Year:       time.Now(),
		BookLangID: "1",
		Deleted:    false,
	},
	{
		ID:         "3",
		Title:      "Third",
		Snippet:    "Third book description",
		GenreID:    "1",
		PagesCnt:   110,
		Year:       time.Now(),
		BookLangID: "1",
		Deleted:    false,
	},
}

type BookReporsitory struct{}

func (b *BookReporsitory) Get(ID string) (models.Book, error) {
	for i := range books {
		if books[i].ID == ID {
			return books[i], nil
		}
	}

	return models.Book{}, fmt.Errorf("Book not found")
}
