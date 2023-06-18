package usecase

import (
	"context"
	"fmt"

	"github.com/hytkgami/trivia-backend/domain"
)

type LobbyInteractor struct {
	LobbyRepository LobbyRepository
}

func (i *LobbyInteractor) CreateLobby(ctx context.Context, uid, name string, public bool) (*domain.Lobby, error) {
	lobby, err := i.LobbyRepository.CreateLobby(ctx, uid, name, public)
	if err != nil {
		return nil, err
	}
	err = i.LobbyRepository.CreateLobbyStatus(ctx, lobby.ID)
	if err != nil {
		return nil, err
	}
	err = i.PublishLobbyStatus(ctx, lobby.ID, domain.LobbyStatusWaiting)
	if err != nil {
		return nil, err
	}
	return lobby, nil
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

func (i *LobbyInteractor) SubscribeLobbyStatus(ctx context.Context, id string, ch chan<- domain.LobbyStatus) error {
	status, err := i.LobbyRepository.FetchLobbyStatus(ctx, id)
	if err != nil {
		return err
	}
	ch <- status
	statusCh := make(chan domain.LobbyStatus)
	i.LobbyRepository.SubscribeLobbyStatus(ctx, id, statusCh)
	go func() {
		for status = range statusCh {
			select {
			case ch <- status:
			default:
				fmt.Println("failed to send lobby status")
				return
			}
		}
	}()
	return nil
}

func (i *LobbyInteractor) PublishLobbyStatus(ctx context.Context, id string, status domain.LobbyStatus) error {
	err := i.LobbyRepository.UpdateLobbyStatus(ctx, id, status)
	if err != nil {
		return err
	}
	return i.LobbyRepository.PublishLobbyStatus(ctx, id, status)
}
