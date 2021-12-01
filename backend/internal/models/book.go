package models

import "time"

type Book struct {
	ID         string
	Title      string
	Snippet    string
	GenreID    string
	PagesCnt   uint
	Year       time.Time
	BookLangID string
	Deleted    bool
}
