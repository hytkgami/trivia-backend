package usecase

import (
	"context"

	"github.com/hytkgami/trivia-backend/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (i *UserInteractor) Signin(ctx context.Context, uid, name string) (*domain.User, error) {
	err := i.UserRepository.UpsertUser(ctx, uid, name)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		UID:  uid,
		Name: name,
	}, nil
}
