package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Book struct {
	ID          string    `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Snippet     string    `json:"snippet" db:"snippet"`
	GenreID     string    `json:"genre_id" db:"genre_id"`
	BookLangID  string    `json:"book_lang_id" db:"book_lang_id"`
	PagesCnt    uint      `json:"pages_cnt" db:"pages_cnt"`
	PublishDate time.Time `json:"pub_date" db:"pub_date"`
	Deleted     bool      `json:"deleted" db:"deleted"`
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
		validation.Field(&b.PublishDate, validation.When(!b.PublishDate.IsZero(), validation.Required)),
		validation.Field(&b.Deleted, validation.In(true, false)),
	)
}
