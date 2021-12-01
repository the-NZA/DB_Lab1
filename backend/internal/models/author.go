package models

import "time"

type Author struct {
	ID        string
	Firstname string
	Lastname  string
	Surname   string
	BirthDate time.Time
	Snippet   string
	Deleted   bool
}
