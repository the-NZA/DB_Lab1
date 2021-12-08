package sqlite3

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/the-NZA/DB_Lab1/backend/internal/models"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storetypes"
)

var (
	insertAuthor = `INSERT INTO authors (firstname, lastname, surname, birth_date, snippet) 
				VALUES (?, ?, ?, ?, ?)`
	updateAuthor = `UPDATE authors
			SET 	firstname = ?, 
				lastname = ?,
				surname = ?,
				birth_date = ?,
				snippet = ?, 
				deleted = ?, 
			WHERE id = ?`
)

type AuthorRepository struct {
	db *sqlx.DB
}

// Get one author from db
func (a *AuthorRepository) Get(ID string) (models.Author, error) {
	author := storetypes.SQLAuthor{}

	err := a.db.Get(&author, "SELECT * FROM authors WHERE id = ? AND deleted != true", ID)
	if err != nil {
		return models.Author{}, err
	}

	return author.ToAuthorModel(), nil
}

// Add one author
func (a *AuthorRepository) Add(author models.Author) (models.Author, error) {
	// Try save new author
	res, err := a.db.Exec(insertAuthor,
		author.Firstname,
		author.Lastname,
		author.Surname,
		author.BirthDate,
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
		author.BirthDate,
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
	var sAuthors []storetypes.SQLAuthor

	// Get all authors from database
	err := a.db.Select(&sAuthors, "SELECT * FROM authors WHERE deleted != true")
	if err != nil {
		return nil, err
	}

	// Convert slice of SQLAuthors to applications Author
	var authors = make([]models.Author, 0, len(sAuthors))
	for i := range sAuthors {
		authors = append(authors, sAuthors[i].ToAuthorModel())
	}

	return authors, nil
}
