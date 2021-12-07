package sqlite3

import (
	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/DB_Lab1/backend/internal/models"
)

type BookRepository struct {
	db *sqlx.DB
}

// Get one book from books
func (b BookRepository) Get(ID string) (models.Book, error) {
	book := models.Book{}

	err := b.db.Get(&book, "SELECT * FROM books WHERE id = ? AND deleted != true", ID)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
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
	return nil
}

// Gell all books from books
func (b BookRepository) GetAll() ([]models.Book, error) {
	return []models.Book{}, nil
}
