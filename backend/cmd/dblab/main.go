package main

import (
	"flag"
	"fmt"

	log "github.com/go-pkgz/lgr"
	"github.com/the-NZA/DB_Lab1/backend/internal/dblab"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "etc/dev.config.json", "Path to config file")
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("[ERROR]: %v", err)
	}
}

// run starts application
func run() error {
	flag.Parse()

	// Create new config
	config := dblab.NewConfig()

	// Read config options
	err := config.ReadFromFile(configPath)
	if err != nil {
		return err
	}

	fmt.Println("Hello! This is app's entrypoint")
	fmt.Printf("%s\n", configPath)

	return nil
}
