package repository

import (
	"context"
	"fmt"

	"github.com/hytkgami/trivia-backend/domain"
	"github.com/jmoiron/sqlx"
)

type ScoreRepository struct {
	DB *sqlx.DB
}

func (r *ScoreRepository) CreateScore(ctx context.Context, score *domain.Score) error {
	fmt.Println(score)
	query := `
		INSERT INTO scores (answer_id, mark, value)
		VALUES (:answer_id, :mark, :value)
		ON DUPLICATE KEY UPDATE mark = VALUES(mark), value = VALUES(value), updated_at = NOW();`
	_, err := r.DB.NamedExecContext(ctx, query, score)
	if err != nil {
		return fmt.Errorf("failed to create a score: %w", err)
	}
	return nil
}

func (r *ScoreRepository) FetchByAnswerIDs(ctx context.Context, answerIDs []string) ([]*domain.Score, error) {
	query := `SELECT answer_id, mark, value FROM scores WHERE answer_id IN (:answer_ids);`
	query, args, err := sqlx.Named(query, map[string]any{"answer_ids": answerIDs})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch scores: %w", err)
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch scores: %w", err)
	}
	query = r.DB.Rebind(query)
	scores := []*domain.Score{}
	err = r.DB.SelectContext(ctx, &scores, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch scores: %w", err)
	}
	return scores, nil
}
