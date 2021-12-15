package sqlite3

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/DB_Lab1/backend/internal/models"
)

var (
	insertBook = `INSERT INTO books (title, snippet, pages_cnt, pub_year, genre_id) 
			VALUES (?, ?, ?, ?, ?)`
	updateBook = `UPDATE books 
			SET 	title = ?, 
				snippet = ?, 
				pages_cnt = ?, 
				pub_year = ?, 
				deleted = ?, 
				genre_id = ? 
			WHERE id = ?`
)

type BookRepository struct {
	db *sqlx.DB
}

// Get one book from books
func (b *BookRepository) Get(ID string) (models.Book, error) {
	book := models.Book{}

	err := b.db.Get(&book, "SELECT * FROM books WHERE id = ? AND deleted != true", ID)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

// Add one book to books
func (b *BookRepository) Add(book models.Book) (models.Book, error) {
	// Try save new book
	res, err := b.db.Exec(insertBook,
		book.Title,
		book.Snippet,
		book.PagesCnt,
		book.PublishYear,
		book.GenreID,
	)
	if err != nil {
		return book, err
	}

	// Try get ID for inserted book
	id, err := res.LastInsertId()
	if err != nil {
		return book, err
	}

	// Save string representation of ID
	book.ID = strconv.FormatInt(id, 10)

	return book, nil
}

// Update one book in books
func (b *BookRepository) Update(book models.Book) (models.Book, error) {
	// Try update book
	_, err := b.db.Exec(updateBook,
		book.Title,
		book.Snippet,
		book.PagesCnt,
		book.PublishYear,
		book.Deleted,
		book.GenreID,
		book.ID,
	)
	if err != nil {
		return book, err
	}

	return book, nil
}

// Delete one book from books
func (b *BookRepository) Delete(ID string) error {
	_, err := b.db.Exec("UPDATE books SET deleted = true WHERE id = ?", ID)

	return err
}

// Gell all books from books
func (b *BookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book

	// Get all books from database
	err := b.db.Select(&books, "SELECT * FROM books WHERE deleted != true")
	if err != nil {
		return nil, err
	}

	return books, nil
}
