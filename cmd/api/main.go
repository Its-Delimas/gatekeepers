package main

import (
	"log"

	"github.com/Its-Delimas/gatekeepers/internals/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	log.Printf("starting server on port %s in %s mode", cfg.Port, cfg.Environment)
}
