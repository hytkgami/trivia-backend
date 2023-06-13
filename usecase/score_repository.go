package usecase

import (
	"context"

	"github.com/hytkgami/trivia-backend/domain"
)

type ScoreRepository interface {
	CreateScore(ctx context.Context, score *domain.Score) error
	FetchByAnswerIDs(ctx context.Context, answerIDs []string) ([]*domain.Score, error)
}
