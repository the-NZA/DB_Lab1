package dblab

import (
	"fmt"
	"net/http"
)

func (s *Server) handleBookGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is get book endpoint")
		book := s.services.BookService().Get()
		fmt.Fprintln(w, book)
	}
}

func (s *Server) handleBookAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is add book endpoint")
		book, err := s.services.BookService().Add()

		if err != nil {
			fmt.Fprintln(w, err)
		}

		fmt.Fprintln(w, book)
	}
}

func (s *Server) handleBookUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is update book endpoint")
		book, err := s.services.BookService().Update()

		if err != nil {
			fmt.Fprintln(w, err)
		}

		fmt.Fprintln(w, book)
	}
}

func (s *Server) handleBookDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is delete book endpoint")
		err := s.services.BookService().Delete()
		if err != nil {
			fmt.Fprintln(w, err)
		}

		fmt.Fprintln(w, "Deleted")
	}
}

func (s *Server) handleBookGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is get all books endpoint")
		books, err := s.services.BookService().GetAll()
		if err != nil {
			fmt.Fprintln(w, err)
		}

		fmt.Fprintln(w, books)
	}
}
