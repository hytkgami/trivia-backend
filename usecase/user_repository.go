package usecase

import (
	"context"

	"github.com/hytkgami/trivia-backend/domain"
)

type UserRepository interface {
	UpsertUser(ctx context.Context, uid string, name string) error
	FetchUsersByUIDs(ctx context.Context, uids []string) ([]*domain.User, error)
}
