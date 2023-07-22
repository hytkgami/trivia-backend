package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"

	"github.com/hytkgami/trivia-backend/graph/model"
	"github.com/hytkgami/trivia-backend/interfaces"
)

// Signin is the resolver for the signin field.
func (r *mutationResolver) Signin(ctx context.Context, name string) (*model.SigninPayload, error) {
	uid, err := interfaces.GetUserUID(ctx)
	if err != nil {
		return nil, err
	}
	user, err := r.UserInteractor.Signin(ctx, uid, name)
	if err != nil {
		return nil, err
	}
	return &model.SigninPayload{
		User: &model.User{
			ID:   user.UID,
			Name: user.Name,
		},
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
