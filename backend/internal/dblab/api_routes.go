package dblab

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/the-NZA/DB_Lab1/backend/internal/models"
)

// handles GET /api/books/:bookID
func (s *Server) handleBookGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is get book endpoint")

		id := chi.URLParam(r, "bookID")
		if id == "" {
			fmt.Fprintln(w, "You must specifiy ID in url param")
			return
		}

		book, err := s.services.BookService().Get(id)
		if err != nil {
			s.error(w, r, 500, err)
			return
		}

		s.respond(w, r, 200, book)
	}
}

// handles POST /api/books/:bookID
func (s *Server) handleBookAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is add book endpoint")

		book, err := s.services.BookService().Add(models.Book{})
		if err != nil {
			s.error(w, r, 500, err)
			return
		}

		s.respond(w, r, 200, book)
	}
}

// handles PUT /api/books/:bookID
func (s *Server) handleBookUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is update book endpoint")

		book, err := s.services.BookService().Update(models.Book{})
		if err != nil {
			s.error(w, r, 500, err)
			return
		}

		s.respond(w, r, 200, book)
	}
}

// handles DELETE /api/books/:bookID
func (s *Server) handleBookDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is delete book endpoint")

		err := s.services.BookService().Delete("")
		if err != nil {
			s.error(w, r, 500, err)
			return
		}

		s.respond(w, r, 200, "Deleted")
	}
}

// handles GET /api/books/all
func (s *Server) handleBookGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is get all books endpoint")

		books, err := s.services.BookService().GetAll()
		if err != nil {
			s.error(w, r, 500, err)
			return
		}

		s.respond(w, r, 200, books)
	}
}
