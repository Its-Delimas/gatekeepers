package main

import (
	"context"
	"log"

	"github.com/Its-Delimas/gatekeepers/internal/config"
	"github.com/Its-Delimas/gatekeepers/internal/db"
	sqlc "github.com/Its-Delimas/gatekeepers/internal/db/sqlc"
	"github.com/Its-Delimas/gatekeepers/internal/user"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	ctx := context.Background()
	pool, err := db.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	userService := user.NewService(sqlc.New(pool))

	newUser, err := userService.Register(ctx, "test2@example.com", "supersecret123")
	if err != nil {
		log.Fatalf("registration failed: %v", err)
	}

	log.Printf("registered user: %s (%s), hash: %s", newUser.ID, newUser.Email, newUser.PasswordHash)
}