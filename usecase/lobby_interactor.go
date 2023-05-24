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
