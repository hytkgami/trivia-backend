package usecase

import (
	"context"
	"fmt"

	"github.com/hytkgami/trivia-backend/domain"
)

type ScoreInteractor struct {
	ScoreRepository ScoreRepository
}

func (i *ScoreInteractor) CreateScore(ctx context.Context, answerID, mark string, value int) error {
	markValue := domain.Mark(mark)
	if !markValue.IsValid() {
		return fmt.Errorf("invalid mark: %s", mark)
	}
	score := &domain.Score{
		AnswerID: answerID,
		Mark:     markValue,
		Value:    value,
	}
	return i.ScoreRepository.CreateScore(ctx, score)
}

func (i *ScoreInteractor) FetchByAnswerIDs(ctx context.Context, answerIDs []string) ([]*domain.Score, error) {
	return i.ScoreRepository.FetchByAnswerIDs(ctx, answerIDs)
}
