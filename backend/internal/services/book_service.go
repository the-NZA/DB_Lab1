package services

import (
	"time"

	"github.com/the-NZA/DB_Lab1/backend/internal/models"
	"github.com/the-NZA/DB_Lab1/backend/internal/store/storer"
)

type BookService struct {
	repository storer.BookReporsitory
}

func (b *BookService) Get() models.Book {
	return models.Book{
		ID:         "1",
		Title:      "First book",
		Snippet:    "This is first book",
		GenreID:    "1",
		PagesCnt:   1,
		Year:       time.Now(),
		BookLangID: "1",
		Deleted:    false,
	}
}

func (b *BookService) Add() {

}

func (b *BookService) Update() {

}

func (b *BookService) Delete() {

}

func (b *BookService) FindAll() {

}
