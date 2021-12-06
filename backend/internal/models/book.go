package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Book struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Snippet    string    `json:"snippet"`
	GenreID    string    `json:"genre_id"`
	BookLangID string    `json:"book_lang_id"`
	PagesCnt   uint      `json:"pages_cnt"`
	Year       time.Time `json:"year"`
	Deleted    bool      `json:"deleted"`
}

// Validate fields which must always have values
func (b Book) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.ID, validation.When(b.ID != "", validation.Required, validation.Length(1, 0))),
		validation.Field(&b.Snippet, validation.When(b.Snippet != "", validation.Required, validation.Length(1, 0))),
		validation.Field(&b.Title, validation.Required, validation.Length(1, 0)),
		validation.Field(&b.GenreID, validation.Required),
		validation.Field(&b.BookLangID, validation.Required),
		validation.Field(&b.PagesCnt, validation.When(b.PagesCnt != 0, validation.Required)),
		validation.Field(&b.Year, validation.When(!b.Year.IsZero(), validation.Required)),
		validation.Field(&b.Deleted, validation.In(true, false)),
	)
}
