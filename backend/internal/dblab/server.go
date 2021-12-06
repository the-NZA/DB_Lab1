package dblab

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-pkgz/lgr"
	"github.com/the-NZA/DB_Lab1/backend/internal/config"
	"github.com/the-NZA/DB_Lab1/backend/internal/services"
)

// Server represents application's main structure
type Server struct {
	config *config.Config
	logger *lgr.Logger
	// store    store.Storer
	services services.Servicer
	router   *chi.Mux
}

// Returns new configured server
func NewServer(c *config.Config, services services.Servicer) (*Server, error) {
	if c == nil {
		return nil, config.ErrEmptyConfig
	}

	return &Server{
		config:   c,
		services: services,
		logger:   configureLogger(c.LogDebug),
		router:   chi.NewRouter(),
	}, nil
}

func (s *Server) configureRouter() {
	s.router.Get("/", s.handleIndexPage())

	s.router.Route("/api", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			s.respond(w, r, http.StatusOK, "This is API route")
		})

		// Books endpoints
		r.Route("/book", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				s.respond(w, r, http.StatusOK, "This is books API route")
			})
			// Get, Add, Update and Delete books by ID
			r.Get("/{bookID}", s.handleBookGet())
			r.Post("/", s.handleBookAdd())
			r.Put("/", s.handleBookUpdate())
			r.Delete("/{bookID}", s.handleBookDelete())
			r.Get("/all", s.handleBookGetAll())
		})

		// Genres endpoints
		r.Route("/genre", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				s.respond(w, r, http.StatusOK, "This is genres API route")
			})

			// Get, Add, Update and Delete books by ID
			r.Get("/{genreID}", s.handleGenreGet())
			r.Post("/", s.handleGenreAdd())
			r.Put("/", s.handleGenreUpdate())
			r.Delete("/{genreID}", s.handleGenreDelete())
			r.Get("/all", s.handleGenreGetAll())
		})

		// Authors endpoints
		r.Route("/author", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				s.respond(w, r, http.StatusOK, "This is authors API route")
			})

			// Get, Add, Update and Delete books by ID
			r.Get("/{authorID}", s.handleAuthorGet())
			r.Post("/", s.handleAuthorAdd())
			r.Put("/", s.handleAuthorUpdate())
			r.Delete("/{authorID}", s.handleAuthorDelete())
			r.Get("/all", s.handleAuthorGetAll())
		})
	})
}

func (s *Server) Start() error {
	s.configureRouter()

	bindAddr := s.config.GetBindAddress()

	s.logger.Logf("[INFO] Server is starting at %v...\n", bindAddr)

	return http.ListenAndServe(bindAddr, s.router)
}

func (s *Server) Shutdown() error {
	return s.services.Close()
}
