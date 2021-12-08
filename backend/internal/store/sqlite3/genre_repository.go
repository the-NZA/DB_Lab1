package sqlite3

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/DB_Lab1/backend/internal/models"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storetypes"
)

var (
	insertGenre = `INSERT INTO books (title, snippet) VALUES (?, ?)`
	updateGenre = `UPDATE books 
			SET 	title = ?, 
				snippet = ?, 
				deleted = ?, 
			WHERE id = ?`
)

type GenreRepository struct {
	db *sqlx.DB
}

// Get one genre
func (g *GenreRepository) Get(ID string) (models.Genre, error) {
	genre := storetypes.SQLGenre{}

	// Try get one genre
	err := g.db.Get(&genre, "SELECT * FROM genres WHERE id = ? AND deleted != true", ID)
	if err != nil {
		return models.Genre{}, err
	}

	return genre.ToGenreModel(), nil
}

// Add new genre
func (g *GenreRepository) Add(genre models.Genre) (models.Genre, error) {
	// Try save new genre
	res, err := g.db.Exec(insertGenre, genre.Title, genre.Snippet)
	if err != nil {
		return genre, err
	}

	// Try get ID for inserted genre
	id, err := res.LastInsertId()
	if err != nil {
		return genre, err
	}

	// Save string representation of ID
	genre.ID = strconv.FormatInt(id, 10)

	return genre, nil
}

// Update genre
func (g *GenreRepository) Update(genre models.Genre) (models.Genre, error) {
	// Try update genre
	_, err := g.db.Exec(updateBook,
		genre.Title,
		genre.Snippet,
		genre.Deleted,
		genre.ID,
	)
	if err != nil {
		return genre, err
	}

	return genre, nil
}

// Delete genre
func (g *GenreRepository) Delete(ID string) error {
	_, err := g.db.Exec("UPDATE genres SET deleted = true WHERE id = ?", ID)

	return err
}

// Get all genres
func (g *GenreRepository) GetAll() ([]models.Genre, error) {
	var sGenres []storetypes.SQLGenre

	// Get all genres from database
	err := g.db.Select(&sGenres, "SELECT * FROM genres WHERE deleted != true")
	if err != nil {
		return nil, err
	}

	// Convert slice of SQLGenre to applications Genre
	var genres = make([]models.Genre, 0, len(sGenres))
	for i := range sGenres {
		genres = append(genres, sGenres[i].ToGenreModel())
	}

	return genres, nil
}
