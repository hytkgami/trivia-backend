package usecase

import (
	"context"

	"github.com/hytkgami/trivia-backend/domain"
)

type QuestionInteractor struct {
	QuestionRepository QuestionRepository
}

type QuestionInput struct {
	Title       string
	OrderNumber int
	Score       int
}

func (i *QuestionInteractor) CreateQuestions(ctx context.Context, uid, lobbyID string, questions []*QuestionInput) ([]*domain.Question, error) {
	return i.QuestionRepository.CreateQuestions(ctx, uid, lobbyID, questions)
}

func (i *QuestionInteractor) FetchQuestionsByLobbyID(ctx context.Context, lobbyID string) ([]*domain.Question, error) {
	return i.QuestionRepository.FetchQuestionsByLobbyID(ctx, lobbyID)
}
