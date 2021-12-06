package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Author struct {
	ID        string
	Firstname string
	Lastname  string
	Surname   string
	BirthDate time.Time
	Snippet   string
	Deleted   bool
}

// Validate fields which must always have values
func (a Author) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field("Firstname", validation.Required, validation.Length(1, 0)),
		validation.Field("Lastname", validation.Required, validation.Length(1, 0)),
		validation.Field("Deleted", validation.Required, validation.In(true, false)),
	)
}
