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
func (b *BookRepository) Add(ba models.BookWithAuthors) (models.BookWithAuthors, error) {
	// Try save new book

	// Start transaction
	tx, err := b.db.Beginx()
	if err != nil {
		return ba, err
	}

	// Remove previous information about book and authors relations
	_, err = tx.Exec("DELETE FROM books_authors WHERE book_id = ?", ba.Book.ID)
	if err != nil {
		tx.Rollback()
		return ba, err
	}

	// Save new book
	res, err := tx.Exec(insertBook,
		ba.Book.Title,
		ba.Book.Snippet,
		ba.Book.PagesCnt,
		ba.Book.PublishYear,
		ba.Book.GenreID,
	)
	if err != nil {
		tx.Rollback()
		return ba, err
	}

	// Try get ID for inserted book
	id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return ba, err
	}

	// Save string representation of ID
	ba.Book.ID = strconv.FormatInt(id, 10)

	// Insert new information about book and authors relations
	for i := range ba.AuthorsIDs {
		_, err = tx.Exec("INSERT INTO books_authors (book_id, author_id) VALUES (?, ?)", ba.Book.ID, ba.AuthorsIDs[i])
		if err != nil {
			tx.Rollback()
			return ba, err
		}
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return ba, err
	}

	// res, err := b.db.Exec(insertBook,
	// 	book.Title,
	// 	book.Snippet,
	// 	book.PagesCnt,
	// 	book.PublishYear,
	// 	book.GenreID,
	// )
	// if err != nil {
	// 	return book, err
	// }

	return ba, nil
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
