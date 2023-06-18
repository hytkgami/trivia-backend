package repository

import (
	"context"
	"fmt"

	"github.com/hytkgami/trivia-backend/domain"
	"github.com/hytkgami/trivia-backend/internal"
	"github.com/hytkgami/trivia-backend/usecase"
	"github.com/jmoiron/sqlx"
)

type LobbyRepository struct {
	DB           *sqlx.DB
	RedisHandler RedisHandler
}

func (r *LobbyRepository) CreateLobby(ctx context.Context, uid, name string, public bool) (*domain.Lobby, error) {
	id, err := internal.NewULID()
	if err != nil {
		return nil, fmt.Errorf("failed to create lobby: %w", err)
	}
	query := `
    INSERT INTO
      lobbies (lobby_id, owner_uid, name, is_public)
    VALUES
      (:lobby_id, :owner_uid, :name, :is_public)
    ;
  `
	_, err = r.DB.NamedExecContext(ctx, query, map[string]any{
		"lobby_id":  id,
		"owner_uid": uid,
		"name":      name,
		"is_public": public,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create lobby: %w", err)
	}
	return &domain.Lobby{
		ID:       id,
		OwnerUID: uid,
		Name:     name,
		IsPublic: public,
	}, nil
}

func (r *LobbyRepository) FetchLobby(ctx context.Context, id string) (*domain.Lobby, error) {
	query := `
		SELECT
			lobby_id, owner_uid, name, is_public
		FROM
			lobbies
		WHERE
			lobby_id = ?
		;
	`
	query = r.DB.Rebind(query)
	var lobby domain.Lobby
	err := r.DB.GetContext(ctx, &lobby, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch lobby: %w", err)
	}
	return &lobby, nil
}

func (r *LobbyRepository) FetchLobbies(ctx context.Context, pagination *usecase.CursorPagination) ([]*domain.Lobby, error) {
	query := `SELECT lobby_id, owner_uid, name, is_public FROM lobbies WHERE is_public = 1`
	cursorQuery, whereArgs, err := generateCursorQuery(pagination)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch lobbies: %w", err)
	}
	query = query + cursorQuery
	if err != nil {
		return nil, fmt.Errorf("failed to fetch lobbies: %w", err)
	}
	query = r.DB.Rebind(query)
	var lobbies []*domain.Lobby
	// FIXME: SelectContextに渡る際にカーソルの絞り込みが無効になっているため修正する
	// 発行SQL: SELECT lobby_id, owner_uid, name, is_public FROM "public"."lobbies" WHERE is_public = TRUE AND $1 > $2 ORDER BY $3 ASC LIMIT $4
	// Args: [lobby_id 01H16KS0GJQ489TMPGMCKJFQW0 lobby_id 4]
	// 他のDBクライアントや、クエリをバインドなしで発行した場合は正常に動作する
	err = r.DB.SelectContext(ctx, &lobbies, query, whereArgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch lobbies: %w", err)
	}
	return lobbies, nil
}

func (r *LobbyRepository) CreateLobbyStatus(ctx context.Context, id string) error {
	query := `
		INSERT INTO
			lobby_lobby_status (lobby_id, lobby_status_id)
		VALUES
			(:lobby_id, :lobby_status_id)
		;
	`
	_, err := r.DB.NamedExecContext(ctx, query, map[string]any{
		"lobby_id":        id,
		"lobby_status_id": domain.LobbyStatusWaitingID,
	})
	return err
}

func (r *LobbyRepository) FetchLobbyStatus(ctx context.Context, id string) (domain.LobbyStatus, error) {
	query := `SELECT ls.status FROM lobby_lobby_status lls JOIN lobby_status ls USING(lobby_status_id) WHERE lls.lobby_id = ?`
	query = r.DB.Rebind(query)
	var status domain.LobbyStatus
	err := r.DB.GetContext(ctx, &status, query, id)
	if err != nil {
		return domain.LobbyStatusWaiting, fmt.Errorf("failed to fetch lobby status: %w", err)
	}
	return status, nil
}

func (r *LobbyRepository) SubscribeLobbyStatus(ctx context.Context, id string, ch chan<- domain.LobbyStatus) {
	channelID := r.lobbyStatusKey(id)
	pubsub := r.RedisHandler.Subscribe(ctx, channelID)
	go func() {
		event := pubsub.Channel()
		for e := range event {
			payload := e.Payload
			select {
			case ch <- domain.LobbyStatus(payload):
			default:
				fmt.Println("failed to send lobby status")
				pubsub.Close()
				return
			}
		}
	}()
}

func (r *LobbyRepository) lobbyStatusKey(id string) string {
	return fmt.Sprintf("lobby:%s:status", id)
}

func (r *LobbyRepository) PublishLobbyStatus(ctx context.Context, id string, status domain.LobbyStatus) error {
	key := r.lobbyStatusKey(id)
	err := r.RedisHandler.Set(ctx, key, status, 0)
	if err != nil {
		return fmt.Errorf("failed to publish lobby status: %w", err)
	}
	err = r.RedisHandler.Publish(ctx, key, status)
	if err != nil {
		return fmt.Errorf("failed to publish lobby status: %w", err)
	}
	return nil
}

func (r *LobbyRepository) UpdateLobbyStatus(ctx context.Context, id string, status domain.LobbyStatus) error {
	query := `UPDATE lobby_lobby_status SET lobby_status_id = ? WHERE lobby_id = ?`
	query = r.DB.Rebind(query)
	_, err := r.DB.ExecContext(ctx, query, status.ID(), id)
	if err != nil {
		return fmt.Errorf("failed to update lobby status: %w", err)
	}
	return nil
}
