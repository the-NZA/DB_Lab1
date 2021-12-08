package sqlite3

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/DB_Lab1/backend/internal/models"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storetypes"
)

var (
	insertBook = `INSERT INTO books (title, snippet, pages_cnt, pub_date, genre_id, book_lang_id) 
			VALUES (?, ?, ?, ?, ?, ?)`
	updateBook = `UPDATE books 
			SET 	title = ?, 
				snippet = ?, 
				pages_cnt = ?, 
				pub_date = ?, 
				deleted = ?, 
				genre_id = ?, 
				book_lang_id = ? 
			WHERE id = ?`
)

type BookRepository struct {
	db *sqlx.DB
}

// Get one book from books
func (b *BookRepository) Get(ID string) (models.Book, error) {
	localBook := storetypes.SQLBook{}

	err := b.db.Get(&localBook, "SELECT * FROM books WHERE id = ? AND deleted != true", ID)
	if err != nil {
		return models.Book{}, err
	}

	return localBook.ToBookModel(), nil
}

// Add one book to books
func (b *BookRepository) Add(book models.Book) (models.Book, error) {
	// Try save new book
	res, err := b.db.Exec(insertBook,
		book.Title,
		book.Snippet,
		book.PagesCnt,
		book.PublishDate,
		book.GenreID,
		book.BookLangID,
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
		book.PublishDate,
		book.Deleted,
		book.GenreID,
		book.BookLangID,
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
