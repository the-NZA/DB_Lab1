package main

import (
	"flag"

	log "github.com/go-pkgz/lgr"
	"github.com/the-NZA/DB_Lab1/backend/internal/config"
	"github.com/the-NZA/DB_Lab1/backend/internal/dblab"
	"github.com/the-NZA/DB_Lab1/backend/internal/services"
	"github.com/the-NZA/DB_Lab1/backend/internal/store"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "etc/dev.config.json", "Path to config file")
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("[ERROR] %v", err)
	}
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

	// Create new server
	server := dblab.NewServer(config, services)

	return server.Start()
}
