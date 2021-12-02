package mock

import (
	"fmt"
	"math/rand"
	"strconv"
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

// Get next ID
func getNextID(b []models.Book) int {
	if len(b) < 1 {
		return 0
	}

	next, err := strconv.Atoi(b[0].ID)
	if err != nil {
		return (rand.Int() % len(b)) + 1
	}

	for i := range b {
		n, err := strconv.Atoi(b[i].ID)
		if err != nil {
			continue
		}

		if n > next {
			next = n
		}
	}

	next++
	return next
}

type BookReporsitory struct{}

// find one book by passed ID
func (b BookReporsitory) findByID(ID string) (int, error) {
	for i := range books {
		if books[i].ID == ID {
			return i, nil
		}
	}

	return -1, fmt.Errorf("Book not found")
}

// Get one book from books
func (b BookReporsitory) Get(ID string) (models.Book, error) {
	idx, err := b.findByID(ID)
	if err != nil {
		return models.Book{}, fmt.Errorf("Book not found")
	}

	return books[idx], nil
}

// Add one book to books
func (b BookReporsitory) Add(book models.Book) (models.Book, error) {
	newID := getNextID(books)

	book.ID = strconv.Itoa(newID)

	books = append(books, book)

	return book, nil
}

// Update one book in books
func (b BookReporsitory) Update(book models.Book) (models.Book, error) {
	idx, err := b.findByID(book.ID)
	if err != nil {
		return book, fmt.Errorf("Book not found")
	}

	books[idx] = book

	return books[idx], nil
}

// Delete one book from books
func (b BookReporsitory) Delete(ID string) error {
	idx, err := b.findByID(ID)
	if err != nil {
		return fmt.Errorf("Book not found")
	}

	books[idx] = books[len(books)-1]
	books = books[:len(books)-1]

	return nil
}
