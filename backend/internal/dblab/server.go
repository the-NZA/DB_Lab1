package dblab

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-pkgz/lgr"
	"github.com/the-NZA/DB_Lab1/backend/internal/config"
	"github.com/the-NZA/DB_Lab1/backend/internal/services"
)

// App represents application's main structure
type App struct {
	config   *config.Config
	logger   *lgr.Logger
	services services.Servicer
	server   *http.Server
	router   *chi.Mux
}

// Returns new configured server
func NewApp(c *config.Config, services services.Servicer) (*App, error) {
	if c == nil {
		return nil, config.ErrEmptyConfig
	}

	return &App{
		config:   c,
		services: services,
		logger:   configureLogger(c.LogDebug),
		router:   chi.NewRouter(),
	}, nil
}

func (a *App) configureRouter() {
	a.router.Get("/", a.handleIndexPage())

	a.router.Route("/api", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			a.respond(w, r, http.StatusOK, "This is API route")
		})

		// Books endpoints
		r.Route("/book", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				a.respond(w, r, http.StatusOK, "This is books API route")
			})
			// Get, Add, Update and Delete books by ID
			r.Get("/{bookID}", a.handleBookGet())
			r.Post("/", a.handleBookAdd())
			r.Put("/", a.handleBookUpdate())
			r.Delete("/{bookID}", a.handleBookDelete())
			r.Get("/all", a.handleBookGetAll())
		})

		// Genres endpoints
		r.Route("/genre", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				a.respond(w, r, http.StatusOK, "This is genres API route")
			})

			// Get, Add, Update and Delete books by ID
			r.Get("/{genreID}", a.handleGenreGet())
			r.Post("/", a.handleGenreAdd())
			r.Put("/", a.handleGenreUpdate())
			r.Delete("/{genreID}", a.handleGenreDelete())
			r.Get("/all", a.handleGenreGetAll())
		})

		// Authors endpoints
		r.Route("/author", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				a.respond(w, r, http.StatusOK, "This is authors API route")
			})

			// Get, Add, Update and Delete books by ID
			r.Get("/{authorID}", a.handleAuthorGet())
			r.Post("/", a.handleAuthorAdd())
			r.Put("/", a.handleAuthorUpdate())
			r.Delete("/{authorID}", a.handleAuthorDelete())
			r.Get("/all", a.handleAuthorGetAll())
		})
	})
}

func (a *App) Start() error {
	a.configureRouter()

	a.server = &http.Server{
		Addr:         a.config.GetBindAddress(),
		Handler:      a.router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	a.logger.Logf("[INFO] Server is starting at %v...\n", a.server.Addr)

	return a.server.ListenAndServe()
}

func (a *App) Shutdown(ctx context.Context) error {
	err := a.services.Close()
	if err != nil {
		return err
	}

	return a.server.Shutdown(ctx)
}
