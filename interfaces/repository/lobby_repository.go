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
	DB *sqlx.DB
}

func (r *LobbyRepository) CreateLobby(ctx context.Context, uid, name string, public bool) (*domain.Lobby, error) {
	id, err := internal.NewULID()
	if err != nil {
		return nil, fmt.Errorf("failed to create lobby: %w", err)
	}
	query := `
    INSERT INTO
      "public"."lobbies" (lobby_id, owner_uid, name, is_public)
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
			"public"."lobbies"
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
	query := `SELECT lobby_id, owner_uid, name, is_public FROM "public"."lobbies" WHERE is_public = TRUE`
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
