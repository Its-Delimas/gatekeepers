package main

import (
	"context"
	"log"

	"github.com/Its-Delimas/gatekeepers/internal/config"
	"github.com/Its-Delimas/gatekeepers/internal/db"
	"github.com/Its-Delimas/gatekeepers/internal/db/sqlc"
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

	queries := sqlc.New(pool)

	user, err := queries.CreateUser(ctx, sqlc.CreateUserParams{
		Email:        "test@gmail.com",
		PasswordHash: "not-a-real-hash-yet",
	})

	if err != nil {
		log.Fatalf("failed to create user: %v", err)
	}
	log.Printf("created user: %s (%s)", user.ID, user.Email)

	// log.Printf("connected to database, starting server on port %s in %s mode", cfg.Port, cfg.Environment)
}
