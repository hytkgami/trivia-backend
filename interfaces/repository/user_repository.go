package repository

import (
	"context"
	"fmt"

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
