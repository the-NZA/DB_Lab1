package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Author struct {
	ID        string    `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Surname   string    `json:"surname"`
	BirthDate time.Time `json:"birth_date"`
	Snippet   string    `json:"snippet"`
	Deleted   bool      `json:"deleted"`
}

// Validate fields which must always have values
func (a Author) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Firstname, validation.Required, validation.Length(1, 0)),
		validation.Field(&a.Lastname, validation.Required, validation.Length(1, 0)),
		validation.Field(&a.Surname, validation.When(a.Surname != "", validation.Required, validation.Length(1, 0))),
		validation.Field(&a.Snippet, validation.When(a.Snippet != "", validation.Required, validation.Length(1, 0))),
		validation.Field(&a.BirthDate, validation.When(!a.BirthDate.IsZero(), validation.Required)),
		validation.Field(&a.Deleted, validation.In(true, false)),
	)
}
