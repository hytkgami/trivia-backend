package usecase

import (
	"context"

	"github.com/hytkgami/trivia-backend/domain"
)

type LobbyInteractor struct {
	LobbyRepository LobbyRepository
}

func (i *LobbyInteractor) CreateLobby(ctx context.Context, uid, name string, public bool) (*domain.Lobby, error) {
	return i.LobbyRepository.CreateLobby(ctx, uid, name, public)
}

func (i *LobbyInteractor) FetchLobby(ctx context.Context, id string) (*domain.Lobby, error) {
	return i.LobbyRepository.FetchLobby(ctx, id)
}

func (i *LobbyInteractor) FetchLobbies(ctx context.Context, params *CursorParams) ([]*domain.Lobby, *PageInfo, error) {
	pagination := params.ToCursorPagination()
	lobbies, err := i.LobbyRepository.FetchLobbies(ctx, pagination)
	if err != nil {
		return nil, nil, err
	}
	pageInfo := &PageInfo{
		HasNextPage: false,
		Cursor:      "",
	}
	if len(lobbies) > pagination.Limit {
		pageInfo.HasNextPage = true
		lobbies = lobbies[:pagination.Limit]
	}
	if len(lobbies) == 0 {
		return lobbies, pageInfo, nil
	}
	last := lobbies[len(lobbies)-1]
	pageInfo.Cursor = last.ID
	return lobbies, pageInfo, nil
}
