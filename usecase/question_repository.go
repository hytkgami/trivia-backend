package usecase

import (
	"context"

	"github.com/hytkgami/trivia-backend/domain"
)

type QuestionRepository interface {
	CreateQuestions(ctx context.Context, uid, lobbyID string, questions []*QuestionInput) ([]*domain.Question, error)
	PublishQuestion(ctx context.Context, lobbyID, questionID string) error
	FetchQuestionByID(ctx context.Context, questionID string) (*domain.Question, error)
	FetchQuestionsByLobbyID(ctx context.Context, lobbyID string) ([]*domain.Question, error)
}
