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
func NewServer(config *config.Config, services services.Servicer) *Server {
	return &Server{
		config:   config,
		services: services,
		// store:  store,
		logger: configureLogger(config.LogDebug),
		router: chi.NewRouter(),
	}
}

func (s *Server) configureRouter() {
	s.router.Get("/", s.handleIndexPage())

	s.router.Route("/api", func(r chi.Router) {
		// Books endpoints
		r.Route("/book", func(r chi.Router) {
			r.Get("/", s.handleBookGet())
		})
	})
}

func (s *Server) Start() error {
	s.configureRouter()

	bindAddr := s.config.GetBindAddress()

	s.logger.Logf("[INFO] Server is starting at %v...\n", bindAddr)

	return http.ListenAndServe(bindAddr, s.router)
}
