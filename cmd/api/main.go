package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/Its-Delimas/gatekeepers/internal/auth"
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

	tokenService := auth.NewTokenService(cfg.JWTSecret)
	userService := user.NewService(sqlc.New(pool))
	userHandler := user.NewHandler(userService, tokenService)

	r := chi.NewRouter()
	r.Post("/register", userHandler.Register)
	r.Post("/login", userHandler.Login)

	log.Printf("listenin on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
