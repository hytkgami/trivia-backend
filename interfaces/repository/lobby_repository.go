package repository

import (
	"context"
	"fmt"

	"github.com/hytkgami/trivia-backend/domain"
	"github.com/hytkgami/trivia-backend/internal"
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
