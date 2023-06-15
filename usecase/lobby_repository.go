package usecase

import (
	"context"

	"github.com/hytkgami/trivia-backend/domain"
)

type LobbyRepository interface {
	CreateLobby(ctx context.Context, uid, name string, public bool) (*domain.Lobby, error)
	FetchLobby(ctx context.Context, id string) (*domain.Lobby, error)
	FetchLobbies(ctx context.Context, pagination *CursorPagination) ([]*domain.Lobby, error)
	CreateLobbyStatus(ctx context.Context, id string) error
	FetchLobbyStatus(ctx context.Context, id string) (domain.LobbyStatus, error)
	SubscribeLobbyStatus(ctx context.Context, id string, ch chan<- domain.LobbyStatus)
	PublishLobbyStatus(ctx context.Context, id string, status domain.LobbyStatus) error
	UpdateLobbyStatus(ctx context.Context, id string, status domain.LobbyStatus) error
}
