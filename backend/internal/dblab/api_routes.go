package dblab

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/the-NZA/DB_Lab1/backend/internal/models"
)

var (
	ErrNoIDSpecified = errors.New("You must specify ID in URL param")
)

/*
* Books endpoints
 */
// handles GET /api/book/:bookID
func (s *Server) handleBookGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "bookID")
		if id == "" {
			s.error(w, r, http.StatusInternalServerError, ErrNoIDSpecified)
			return
		}

		book, err := s.services.BookService().Get(id)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, book)
	}
}

// handles POST /api/book/:bookID
func (s *Server) handleBookAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		book, err := s.services.BookService().Add(models.Book{})
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusCreated, book)
	}
}

// handles PUT /api/book/:bookID
func (s *Server) handleBookUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		book, err := s.services.BookService().Update(models.Book{})
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, book)
	}
}

// handles DELETE /api/book/:bookID
func (s *Server) handleBookDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := s.services.BookService().Delete("")
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, "Deleted")
	}
}

// handles GET /api/book/all
func (s *Server) handleBookGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := s.services.BookService().GetAll()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, books)
	}
}

/*
* Genres endpoints
 */
// handles GET /api/genre/:genreID
func (s *Server) handleGenreGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "genreID")
		if id == "" {
			s.error(w, r, http.StatusInternalServerError, ErrNoIDSpecified)
			return
		}

		genre, err := s.services.GenreService().Get(id)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, genre)
	}
}

// handles POST /api/genre/:genreID
func (s *Server) handleGenreAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		genre, err := s.services.GenreService().Add(models.Genre{})
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusCreated, genre)
	}
}

// handles PUT /api/genre/:genreID
func (s *Server) handleGenreUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		genre, err := s.services.GenreService().Update(models.Genre{})
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, genre)
	}
}

// handles DELETE /api/genre/:genreID
func (s *Server) handleGenreDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := s.services.GenreService().Delete("")
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, "Deleted")
	}
}

// handles GET /api/genre/all
func (s *Server) handleGenreGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		genres, err := s.services.GenreService().GetAll()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, genres)
	}
}
