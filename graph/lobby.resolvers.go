package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"strings"

	"github.com/hytkgami/trivia-backend/domain"
	"github.com/hytkgami/trivia-backend/graph/helper"
	"github.com/hytkgami/trivia-backend/graph/loader"
	"github.com/hytkgami/trivia-backend/graph/model"
	"github.com/hytkgami/trivia-backend/interfaces"
	"github.com/hytkgami/trivia-backend/usecase"
)

// Owner is the resolver for the owner field.
func (r *lobbyResolver) Owner(ctx context.Context, obj *model.Lobby) (*model.User, error) {
	user, err := loader.LoadUser(ctx, obj.OwnerUID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Questions is the resolver for the questions field.
func (r *lobbyResolver) Questions(ctx context.Context, obj *model.Lobby) ([]*model.Question, error) {
	questions, err := r.QuestionInteractor.FetchQuestionsByLobbyID(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Question, len(questions))
	for i, q := range questions {
		result[i] = &model.Question{
			ID:          q.ID,
			Title:       q.Title,
			OrderNumber: q.OrderNumber,
			Score:       q.Score,
		}
	}
	return result, nil
}

// PublishLobbyStatus is the resolver for the publishLobbyStatus field.
func (r *mutationResolver) PublishLobbyStatus(ctx context.Context, lobbyID string, status model.LobbyStatus) (*model.PublishLobbyStatusPayload, error) {
	s := domain.LobbyStatus(strings.ToLower(string(status)))
	if !s.Valid() {
		return nil, fmt.Errorf("invalid lobby status: %s", status)
	}
	err := r.LobbyInteractor.PublishLobbyStatus(ctx, lobbyID, s)
	if err != nil {
		return nil, err
	}
	lobby, err := r.LobbyInteractor.FetchLobby(ctx, lobbyID)
	if err != nil {
		return nil, err
	}
	return &model.PublishLobbyStatusPayload{
		Lobby: &model.Lobby{
			ID:       lobby.ID,
			Name:     lobby.Name,
			Public:   lobby.IsPublic,
			OwnerUID: lobby.OwnerUID,
		},
		Status: status,
	}, nil
}

// CreateLobby is the resolver for the createLobby field.
func (r *mutationResolver) CreateLobby(ctx context.Context, name string, public bool) (*model.CreateLobbyPayload, error) {
	uid, err := interfaces.GetUserUID(ctx)
	if err != nil {
		return nil, err
	}
	lobby, err := r.LobbyInteractor.CreateLobby(ctx, uid, name, public)
	if err != nil {
		return nil, err
	}
	return &model.CreateLobbyPayload{
		Lobby: &model.Lobby{
			ID:     lobby.ID,
			Name:   lobby.Name,
			Public: lobby.IsPublic,
		},
	}, nil
}

// DeleteLobby is the resolver for the deleteLobby field.
func (r *mutationResolver) DeleteLobby(ctx context.Context, id string) (*model.Lobby, error) {
	panic(fmt.Errorf("not implemented: DeleteLobby - deleteLobby"))
}

// Lobbies is the resolver for the lobbies field.
func (r *queryResolver) Lobbies(ctx context.Context, first *int, last *int, after *string, before *string, orderDirection model.OrderDirection, orderBy model.LobbiesQueryOrderBy) (*model.LobbyConnection, error) {
	if err := helper.ValidateRelayCursor(first, after, last, before); err != nil {
		return nil, err
	}
	params := &usecase.CursorParams{
		First:   first,
		Last:    last,
		After:   after,
		Before:  before,
		OrderBy: strings.ToLower(orderBy.String()),
	}
	lobbies, pageInfo, err := r.LobbyInteractor.FetchLobbies(ctx, params)
	if err != nil {
		return nil, err
	}
	edges := make([]*model.LobbyEdge, len(lobbies))
	for i, lobby := range lobbies {
		edges[i] = &model.LobbyEdge{
			Node: &model.Lobby{
				ID:       lobby.ID,
				Name:     lobby.Name,
				Public:   lobby.IsPublic,
				OwnerUID: lobby.OwnerUID,
			},
			Cursor: lobby.ID,
		}
	}
	return &model.LobbyConnection{
		Edges: edges,
		PageInfo: &model.PageInfo{
			HasNextPage: pageInfo.HasNextPage,
			Cursor:      pageInfo.Cursor,
		},
	}, nil
}

// Lobby is the resolver for the lobby field.
func (r *queryResolver) Lobby(ctx context.Context, id string) (*model.Lobby, error) {
	lobby, err := r.LobbyInteractor.FetchLobby(ctx, id)
	if err != nil {
		return nil, err
	}
	return &model.Lobby{
		ID:       lobby.ID,
		Name:     lobby.Name,
		Public:   lobby.IsPublic,
		OwnerUID: lobby.OwnerUID,
	}, nil
}

// LobbyStatus is the resolver for the lobbyStatus field.
func (r *subscriptionResolver) LobbyStatus(ctx context.Context, lobbyID string) (<-chan model.LobbyStatus, error) {
	ch := make(chan domain.LobbyStatus)
	go func() {
		err := r.LobbyInteractor.SubscribeLobbyStatus(ctx, lobbyID, ch)
		if err != nil {
			fmt.Println(err)
			return
		}
	}()
	modelCh := make(chan model.LobbyStatus)
	go func() {
		for s := range ch {
			m := model.LobbyStatus(strings.ToUpper(string(s)))
			select {
			case modelCh <- m:
				fmt.Println("sent lobby status")
			default:
				fmt.Println("channel closed")
				return
			}
		}
	}()
	return modelCh, nil
}

// Lobby returns LobbyResolver implementation.
func (r *Resolver) Lobby() LobbyResolver { return &lobbyResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type lobbyResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
