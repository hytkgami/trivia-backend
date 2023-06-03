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
	DB           *sqlx.DB
	RedisHandler RedisHandler
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

func (r *QuestionRepository) PublishQuestion(ctx context.Context, lobbyID, questionID string) error {
	key := fmt.Sprintf("lobby:%s:question", lobbyID)
	err := r.RedisHandler.Set(ctx, key, questionID, 0)
	if err != nil {
		return fmt.Errorf("failed to publish question: %w", err)
	}
	return nil
}

func (r *QuestionRepository) FetchQuestionByID(ctx context.Context, questionID string) (*domain.Question, error) {
	query := `
		SELECT question_id, created_by, lobby_id, title, order_number, score
		FROM questions
		WHERE question_id = $1
		LIMIT 1;
	`
	var question domain.Question
	err := r.DB.GetContext(ctx, &question, query, questionID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch question by id: %w", err)
	}
	return &question, nil
}

func (r *QuestionRepository) FetchQuestionsByLobbyID(ctx context.Context, lobbyID string) ([]*domain.Question, error) {
	query := `
		SELECT question_id, created_by, lobby_id, title, order_number, score
		FROM questions
		WHERE lobby_id = $1
		ORDER BY order_number ASC
	`
	var dbQuestions []*domain.Question
	err := r.DB.SelectContext(ctx, &dbQuestions, query, lobbyID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch questions by lobby id: %w", err)
	}
	return dbQuestions, nil
}
