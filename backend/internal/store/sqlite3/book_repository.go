package sqlite3

import (
	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/DB_Lab1/backend/internal/models"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storetypes"
)

type BookRepository struct {
	db *sqlx.DB
}

// Get one book from books
func (b BookRepository) Get(ID string) (models.Book, error) {
	localBook := storetypes.SQLBook{}

	err := b.db.Get(&localBook, "SELECT * FROM books WHERE id = ? AND deleted != true", ID)
	if err != nil {
		return models.Book{}, err
	}

	return localBook.ToBookModel(), nil
}

// Add one book to books
func (b BookRepository) Add(book models.Book) (models.Book, error) {
	return models.Book{}, nil
}

// Update one book in books
func (b BookRepository) Update(book models.Book) (models.Book, error) {
	return models.Book{}, nil
}

// Delete one book from books
func (b BookRepository) Delete(ID string) error {
	_, err := b.db.Exec("UPDATE books SET deleted = true WHERE id = ?", ID)

	return err
}

// Gell all books from books
func (b BookRepository) GetAll() ([]models.Book, error) {
	var sBooks []storetypes.SQLBook

	// Get all books from database
	err := b.db.Select(&sBooks, "SELECT * FROM books WHERE deleted != true")
	if err != nil {
		return nil, err
	}

	// Convert slice of SQLBooks to applications books
	var books = make([]models.Book, 0, len(sBooks))
	for i := range sBooks {
		books = append(books, sBooks[i].ToBookModel())
	}

	return books, nil
}
