package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Book struct {
	ID         string
	Title      string
	Snippet    string
	GenreID    string
	BookLangID string
	PagesCnt   uint
	Year       time.Time
	Deleted    bool
}

// Validate fields which must always have values
func (b Book) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field("Title", validation.Required, validation.Length(1, 0)),
		validation.Field("GenreID", validation.Required),
		validation.Field("BookLangID", validation.Required),
		validation.Field("Deleted", validation.Required, validation.In(true, false)),
	)
}
