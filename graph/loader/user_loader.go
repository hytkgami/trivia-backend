package loader

import (
	"context"
	"errors"

	"github.com/graph-gophers/dataloader/v7"
	"github.com/hytkgami/trivia-backend/graph/model"
	"github.com/hytkgami/trivia-backend/usecase"
)

type UserLoader struct {
	Interactor *usecase.UserInteractor
}

func (l *UserLoader) BatchGet(ctx context.Context, uids []string) []*dataloader.Result[*model.User] {
	result := make([]*dataloader.Result[*model.User], len(uids))
	users, err := l.Interactor.FetchUsersByUIDs(ctx, uids)
	if err != nil {
		for i := range uids {
			result[i] = &dataloader.Result[*model.User]{Data: nil, Error: err}
		}
		return result
	}
	mapped := map[string]*model.User{}
	for _, u := range users {
		mapped[u.UID] = &model.User{
			ID:   u.UID,
			Name: u.Name,
		}
	}
	for i, uid := range uids {
		owner, ok := mapped[uid]
		if !ok {
			err := errors.New("user not found")
			result[i] = &dataloader.Result[*model.User]{Data: nil, Error: err}
		}
		result[i] = &dataloader.Result[*model.User]{Data: owner, Error: nil}
	}
	return result
}

func LoadUser(ctx context.Context, uid string) (*model.User, error) {
	loaders := GetLoaders(ctx)
	thunk := loaders.UserLoader.Load(ctx, uid)
	return thunk()
}
