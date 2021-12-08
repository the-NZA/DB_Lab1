package sqlite3

import (
	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/DB_Lab1/backend/internal/models"
)

var (
	insertGenre = `INSERT INTO books (title, snippet, pages_cnt, pub_date, genre_id, book_lang_id) 
			VALUES (?, ?, ?, ?, ?, ?)`
	updateGenre = `UPDATE books 
			SET 	title = ?, 
				snippet = ?, 
				pages_cnt = ?, 
				pub_date = ?, 
				deleted = ?, 
				genre_id = ?, 
				book_lang_id = ? 
			WHERE id = ?`
)

type GenreRepository struct {
	db *sqlx.DB
}

// Get one genre
func (g *GenreRepository) Get(ID string) (models.Genre, error) {
	return models.Genre{}, nil
}

// Add new genre
func (g *GenreRepository) Add(genre models.Genre) (models.Genre, error) {
	return models.Genre{}, nil
}

// Update genre
func (g *GenreRepository) Update(genre models.Genre) (models.Genre, error) {
	return models.Genre{}, nil
}

// Delete genre
func (g *GenreRepository) Delete(ID string) error {
	return nil
}

// Get all genres
func (g *GenreRepository) GetAll() ([]models.Genre, error) {
	return nil, nil
}
