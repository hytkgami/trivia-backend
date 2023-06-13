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
		ON CONFLICT (question_id, uid) DO UPDATE SET content = :content, updated_at = NOW();
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

func (r *AnswerRepository) FetchByQuestionIDs(ctx context.Context, questionIDs []string) ([]*domain.Answer, error) {
	query := `
		SELECT answer_id, uid, question_id, content FROM answers
		WHERE question_id IN (:question_ids)
	`
	query, args, err := sqlx.Named(query, map[string]any{
		"question_ids": questionIDs,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch answers: %w", err)
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch answers: %w", err)
	}
	query = r.DB.Rebind(query)
	var answers []*domain.Answer
	err = r.DB.SelectContext(ctx, &answers, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch answers: %w", err)
	}
	return answers, nil
}

func (r *AnswerRepository) FetchByID(ctx context.Context, id string) (*domain.Answer, error) {
	query := "SELECT answer_id, uid, question_id, content FROM answers WHERE answer_id = $1"
	var answer domain.Answer
	err := r.DB.GetContext(ctx, &answer, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch an answer: %w", err)
	}
	return &answer, nil
}

func (r *AnswerRepository) FetchByIDs(ctx context.Context, ids []string) ([]*domain.Answer, error) {
	query := "SELECT answer_id, uid, question_id, content FROM answers WHERE answer_id IN (:ids)"
	query, args, err := sqlx.Named(query, map[string]any{
		"answer_ids": ids,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch answers: %w", err)
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch answers: %w", err)
	}
	query = r.DB.Rebind(query)
	var answers []*domain.Answer
	err = r.DB.SelectContext(ctx, &answers, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch answers: %w", err)
	}
	return answers, nil
}
