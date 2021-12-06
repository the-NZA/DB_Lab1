package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/go-pkgz/lgr"
	"github.com/the-NZA/DB_Lab1/backend/internal/config"
	"github.com/the-NZA/DB_Lab1/backend/internal/dblab"
	"github.com/the-NZA/DB_Lab1/backend/internal/services"
	"github.com/the-NZA/DB_Lab1/backend/internal/store"
)

var (
	configPath string
	done       = make(chan struct{}, 1)
	quit       = make(chan os.Signal, 1)
)

func init() {
	flag.StringVar(&configPath, "config", "etc/dev.config.json", "Path to config file")
}

func main() {
	if err := run(); err != http.ErrServerClosed {
		log.Fatalf("[ERROR] %v", err)
	}

	<-done
	log.Printf("[INFO] Server is down")
}

// run starts application
func run() error {
	flag.Parse()

	// Create new config
	config := config.NewConfig()

	// Read config options
	err := config.ReadFromFile(configPath)
	if err != nil {
		return err
	}

	// Create new store base on passed config
	store, err := store.NewStore(config)
	if err != nil {
		return err
	}

	// Init services with created store
	services, err := services.NewServices(config, store)
	if err != nil {
		return err
	}

	// Create new app
	app, err := dblab.NewApp(config, services)
	if err != nil {
		return err
	}

	// Register signal listener
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Run shutdown goroutin
	go shutdown(app, quit, done)

	return app.Start()
}

// shutdown performs app's killing action
func shutdown(a *dblab.App, quit <-chan os.Signal, done chan<- struct{}) {
	<-quit

	log.Printf("[INFO] Server is shutting down...\n")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := a.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}

	close(done)
}
