package storetypes

import (
	"database/sql"

	"github.com/the-NZA/DB_Lab1/backend/internal/models"
)

// SQLBook type add NULL values support
type SQLBook struct {
	ID          string         `db:"id"`
	Title       string         `db:"title"`
	Snippet     sql.NullString `db:"snippet"`
	GenreID     string         `db:"genre_id"`
	BookLangID  string         `db:"book_lang_id"`
	PagesCnt    sql.NullInt32  `db:"pages_cnt"`
	PublishDate sql.NullTime   `db:"pub_date"`
	Deleted     bool           `db:"deleted"`
}

// Convert SQLBook to application's Book model
func (sb *SQLBook) ToBookModel() models.Book {
	book := models.Book{
		ID:         sb.ID,
		Title:      sb.Title,
		GenreID:    sb.GenreID,
		BookLangID: sb.BookLangID,
		Deleted:    sb.Deleted,
	}

	if sb.Snippet.Valid {
		book.Snippet = sb.Snippet.String
	}

	if sb.PagesCnt.Valid {
		book.PagesCnt = uint(sb.PagesCnt.Int32)
	}

	if sb.PublishDate.Valid {
		book.PublishDate = sb.PublishDate.Time
	}

	return book
}

// SQLAuthor type add NULL values support
type SQLAuthor struct {
	ID        string         `db:"id"`
	Firstname string         `db:"firstname"`
	Lastname  string         `db:"lastname"`
	Surname   sql.NullString `db:"surname"`
	BirthDate sql.NullTime   `db:"birth_date"`
	Snippet   sql.NullString `db:"snippet"`
	Deleted   bool           `db:"deleted"`
}

// Convert SQLBook to application's Author model
func (au *SQLAuthor) ToAuthorModel() models.Author {
	author := models.Author{
		ID:        au.ID,
		Firstname: au.Firstname,
		Lastname:  au.Lastname,
		Deleted:   au.Deleted,
	}

	if au.Snippet.Valid {
		author.Snippet = au.Snippet.String
	}

	if au.BirthDate.Valid {
		author.BirthDate = au.BirthDate.Time
	}

	if au.Surname.Valid {
		author.Snippet = au.Snippet.String
	}

	return author
}

// SQLGenre type add NULL values support
type SQLGenre struct {
	ID      string         `db:"id"`
	Title   string         `db:"title"`
	Snippet sql.NullString `db:"snippet"`
	Deleted bool           `db:"deleted"`
}

// Convert SQLGenre to application's Genre model
func (g *SQLGenre) ToGenreModel() models.Genre {
	genre := models.Genre{
		ID:      g.ID,
		Title:   g.Title,
		Deleted: g.Deleted,
	}

	if g.Snippet.Valid {
		genre.Snippet = g.Snippet.String
	}

	return genre
}
