package repository

import (
	"context"
	"fmt"

	"github.com/hytkgami/trivia-backend/domain"
	"github.com/hytkgami/trivia-backend/internal"
	"github.com/hytkgami/trivia-backend/usecase"
	"github.com/jmoiron/sqlx"
)

type QuestionRepository struct {
	DB *sqlx.DB
}

func (r *QuestionRepository) CreateQuestions(ctx context.Context, uid, lobbyID string, questions []*usecase.QuestionInput) ([]*domain.Question, error) {
	query := `
    INSERT INTO questions (question_id, created_by, lobby_id, title, order_number, score)
    VALUES (:question_id, :created_by, :lobby_id, :title, :order_number, :score)
  `
	dbQuestions := make([]*domain.Question, len(questions))
	for i, q := range questions {
		id, err := internal.NewULID()
		if err != nil {
			return nil, fmt.Errorf("failed to create questions: %w", err)
		}
		dbQuestions[i] = &domain.Question{
			ID:          id,
			CreatedBy:   uid,
			LobbyID:     lobbyID,
			Title:       q.Title,
			OrderNumber: q.OrderNumber,
			Score:       q.Score,
		}
	}
	_, err := r.DB.NamedExecContext(ctx, query, dbQuestions)
	if err != nil {
		return nil, fmt.Errorf("failed to create questions: %w", err)
	}
	return dbQuestions, nil
}
