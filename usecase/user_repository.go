package usecase

import (
	"context"
)

type UserRepository interface {
	UpsertUser(ctx context.Context, uid string, name string) error
}
