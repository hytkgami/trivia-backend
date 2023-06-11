package repository

import (
	"context"
	"fmt"

	"github.com/hytkgami/trivia-backend/domain"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func (r *UserRepository) UpsertUser(ctx context.Context, uid string, name string) error {
	query := `
    INSERT INTO
      users (uid, name)
    VALUES
      (:uid, :name)
    ON CONFLICT (uid) DO UPDATE SET name = :name, updated_at = NOW();
  `
	_, err := r.DB.NamedExecContext(ctx, query, map[string]any{
		"uid":  uid,
		"name": name,
	})
	if err != nil {
		return fmt.Errorf("failed to upsert user: %w", err)
	}
	return nil
}

func (r *UserRepository) FetchUsersByUIDs(ctx context.Context, uids []string) ([]*domain.User, error) {
	query, args, err := sqlx.In(`
		SELECT
			uid, name
		FROM
			users
		WHERE
			uid IN (?)
	`, uids)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}
	query = r.DB.Rebind(query)
	var users []*domain.User
	err = r.DB.SelectContext(ctx, &users, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}
	return users, nil
}
