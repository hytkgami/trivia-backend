package usecase

import (
	"context"

	"github.com/hytkgami/trivia-backend/domain"
)

type AnswerRepository interface {
	Create(ctx context.Context, uid, quesitonID, content string) (*domain.Answer, error)
}
