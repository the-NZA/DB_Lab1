package services

import "github.com/the-NZA/DB_Lab1/backend/internal/store/storer"

type BookService struct {
	repository storer.BookReporsitory
}

func (b *BookService) Get() string {
	return "Book is received"
}

func (b *BookService) Add() {

}

func (b *BookService) Update() {

}

func (b *BookService) Delete() {

}

func (b *BookService) FindAll() {

}
