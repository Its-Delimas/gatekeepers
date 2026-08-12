package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/Its-Delimas/gatekeepers/internal/db/sqlc"
	"golang.org/x/crypto/bcrypt"
)

var ErrEmailTaken = errors.New("email is already registered")

type Service struct {
	queries *sqlc.Queries
}

func NewService(queries *sqlc.Queries) *Service {
	return &Service{queries: queries}
}

func (s *Service) Register(ctx context.Context, email, password string) (sqlc.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return sqlc.User{}, fmt.Errorf("failed to hash password: %w", err)
	}
	user, err := s.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Email:        email,
		PasswordHash: string(hash),
	})
	if err != nil {
		return sqlc.User{},fmt.Errorf("failed to create user: %w",err)
	}
	return user,nil
}
