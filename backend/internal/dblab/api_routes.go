package dblab

import (
	"encoding/json"
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
func (s *App) handleBookGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "bookID")
		if id == "" {
			s.error(w, r, http.StatusInternalServerError, ErrNoIDSpecified)
			return
		}

		// Try get book
		book, err := s.services.BookService().Get(id)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, book)
	}
}

// handles POST /api/book
func (s *App) handleBookAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			book models.Book
			err  error
		)

		if err = json.NewDecoder(r.Body).Decode(&book); err != nil {
			s.logger.Logf("[INFO] During body parse: %v\n", err)
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		// Try save new book
		book, err = s.services.BookService().Add(book)
		if err != nil {
			s.logger.Logf("[INFO] During book saving: %v\n", err)
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusCreated, book)
	}
}

// handles PUT /api/book
func (s *App) handleBookUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			book models.Book
			err  error
		)

		if err = json.NewDecoder(r.Body).Decode(&book); err != nil {
			s.logger.Logf("[INFO] During body parse: %v\n", err)
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		// Try update book
		book, err = s.services.BookService().Update(book)
		if err != nil {
			s.logger.Logf("[INFO] During book updating: %v\n", err)
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, book)
	}
}

// handles DELETE /api/book/:bookID
func (s *App) handleBookDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "bookID")
		if id == "" {
			s.error(w, r, http.StatusInternalServerError, ErrNoIDSpecified)
			return
		}

		// Try delete book by ID
		err := s.services.BookService().Delete(id)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, "Deleted")
	}
}

// handles GET /api/book/all
func (s *App) handleBookGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Try get all books
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
func (a *App) handleGenreGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "genreID")
		if id == "" {
			a.error(w, r, http.StatusInternalServerError, ErrNoIDSpecified)
			return
		}

		genre, err := a.services.GenreService().Get(id)
		if err != nil {
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusOK, genre)
	}
}

// handles POST /api/genre
func (a *App) handleGenreAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		genre, err := a.services.GenreService().Add(models.Genre{})
		if err != nil {
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusCreated, genre)
	}
}

// handles PUT /api/genre
func (a *App) handleGenreUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		genre, err := a.services.GenreService().Update(models.Genre{})
		if err != nil {
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusOK, genre)
	}
}

// handles DELETE /api/genre/:genreID
func (a *App) handleGenreDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := a.services.GenreService().Delete("")
		if err != nil {
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusOK, "Deleted")
	}
}

// handles GET /api/genre/all
func (a *App) handleGenreGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		genres, err := a.services.GenreService().GetAll()
		if err != nil {
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusOK, genres)
	}
}

/*
* Authors endpoints
 */
// handles GET /api/author/:authorID
func (a *App) handleAuthorGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "authorID")
		if id == "" {
			a.error(w, r, http.StatusInternalServerError, ErrNoIDSpecified)
			return
		}

		author, err := a.services.AuthorService().Get(id)
		if err != nil {
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusOK, author)
	}
}

// handles POST /api/author
func (a *App) handleAuthorAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author, err := a.services.AuthorService().Add(models.Author{})
		if err != nil {
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusCreated, author)
	}
}

// handles PUT /api/author
func (a *App) handleAuthorUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author, err := a.services.AuthorService().Update(models.Author{})
		if err != nil {
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusOK, author)
	}
}

// handles DELETE /api/author/:authorID
func (a *App) handleAuthorDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := a.services.AuthorService().Delete("")
		if err != nil {
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusOK, "Deleted")
	}
}

// handles GET /api/author/all
func (a *App) handleAuthorGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authors, err := a.services.AuthorService().GetAll()
		if err != nil {
			a.error(w, r, http.StatusInternalServerError, err)
			return
		}

		a.respond(w, r, http.StatusOK, authors)
	}
}
