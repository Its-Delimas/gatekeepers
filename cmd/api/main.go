package main

import (
	"context"
	"log"

	"github.com/Its-Delimas/gatekeepers/internals/config"
	"github.com/Its-Delimas/gatekeepers/internals/db"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	ctx := context.Background()
	pool, err := db.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	defer pool.Close()

	log.Printf("connected to database, starting server on port %s in %s mode", cfg.Port, cfg.Environment)
}
