package usecase

import (
	"context"

	"github.com/hytkgami/trivia-backend/domain"
)

type LobbyRepository interface {
	CreateLobby(ctx context.Context, uid, name string, public bool) (*domain.Lobby, error)
	FetchLobby(ctx context.Context, id string) (*domain.Lobby, error)
	FetchLobbies(ctx context.Context, pagination *CursorPagination) ([]*domain.Lobby, error)
}
