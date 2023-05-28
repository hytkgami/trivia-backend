package usecase

import (
	"context"

	"github.com/hytkgami/trivia-backend/domain"
)

type AnswerInteractor struct {
	AnswerRepository AnswerRepository
}

func (i *AnswerInteractor) Create(ctx context.Context, uid, questionID, content string) (*domain.Answer, error) {
	return i.AnswerRepository.Create(ctx, uid, questionID, content)
}
