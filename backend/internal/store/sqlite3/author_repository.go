package sqlite3

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/DB_Lab1/backend/internal/models"
)

var (
	insertAuthor = `INSERT INTO authors (firstname, lastname, surname,  snippet) 
				VALUES (?, ?, ?, ?, ?)`
	updateAuthor = `UPDATE authors
			SET 	firstname = ?, 
				lastname = ?,
				surname = ?,
				snippet = ?, 
				deleted = ?, 
			WHERE id = ?`
)

type AuthorRepository struct {
	db *sqlx.DB
}

// Get one author from db
func (a *AuthorRepository) Get(ID string) (models.Author, error) {
	author := models.Author{}

	err := a.db.Get(&author, "SELECT * FROM authors WHERE id = ? AND deleted != true", ID)
	if err != nil {
		return models.Author{}, err
	}

	return author, nil
}

// Add one author
func (a *AuthorRepository) Add(author models.Author) (models.Author, error) {
	// Try save new author
	res, err := a.db.Exec(insertAuthor,
		author.Firstname,
		author.Lastname,
		author.Surname,
		author.Snippet,
	)
	if err != nil {
		return author, err
	}

	// Try get ID for inserted author
	id, err := res.LastInsertId()
	if err != nil {
		return author, err
	}

	// Save string representation of ID
	author.ID = strconv.FormatInt(id, 10)

	return author, nil
}

// Update one author
func (a *AuthorRepository) Update(author models.Author) (models.Author, error) {
	// Try update author
	_, err := a.db.Exec(updateAuthor,
		author.Firstname,
		author.Lastname,
		author.Surname,
		author.Snippet,
		author.Deleted,
		author.ID,
	)
	if err != nil {
		return author, err
	}

	return author, nil
}

// Delete one author
func (a *AuthorRepository) Delete(ID string) error {
	_, err := a.db.Exec("UPDATE authors SET deleted = true WHERE id = ?", ID)

	return err
}

// Gell all authors
func (a *AuthorRepository) GetAll() ([]models.Author, error) {
	var authors []models.Author

	// Get all authors from database
	err := a.db.Select(&authors, "SELECT * FROM authors WHERE deleted != true")
	if err != nil {
		return nil, err
	}

	return authors, nil
}
