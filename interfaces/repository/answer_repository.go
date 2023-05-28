package repository

import (
	"context"
	"fmt"

	"github.com/hytkgami/trivia-backend/domain"
	"github.com/hytkgami/trivia-backend/internal"
	"github.com/jmoiron/sqlx"
)

type AnswerRepository struct {
	DB *sqlx.DB
}

func (r *AnswerRepository) Create(ctx context.Context, uid, quesitonID, content string) (*domain.Answer, error) {
	query := `
    INSERT INTO answers (answer_id, uid, question_id, content)
    VALUES (:answer_id, :uid, :question_id, :content)
  `
	id, err := internal.NewULID()
	if err != nil {
		return nil, fmt.Errorf("failed to create an answer: %w", err)
	}
	answer := &domain.Answer{
		ID:         id,
		UID:        uid,
		QuestionID: quesitonID,
		Content:    content,
	}
	_, err = r.DB.NamedExecContext(ctx, query, answer)
	if err != nil {
		return nil, fmt.Errorf("failed to create an answer: %w", err)
	}
	return answer, nil
}
