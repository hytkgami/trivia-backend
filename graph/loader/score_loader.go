package loader

import (
	"context"
	"strings"

	"github.com/graph-gophers/dataloader/v7"
	"github.com/hytkgami/trivia-backend/graph/model"
	"github.com/hytkgami/trivia-backend/usecase"
)

type ScoreLoader struct {
	ScoreInteractor *usecase.ScoreInteractor
}

func (l *ScoreLoader) BatchGetByAnswerIDs(ctx context.Context, answerIDs []string) []*dataloader.Result[*model.Score] {
	result := make([]*dataloader.Result[*model.Score], len(answerIDs))
	scores, err := l.ScoreInteractor.FetchByAnswerIDs(ctx, answerIDs)
	if err != nil {
		for i := range answerIDs {
			result[i] = &dataloader.Result[*model.Score]{Data: nil, Error: err}
		}
		return result
	}
	scoreMap := map[string]*model.Score{}
	for _, score := range scores {
		mark := model.Mark(strings.ToUpper(score.Mark.String()))
		scoreMap[score.AnswerID] = &model.Score{
			Mark:  mark,
			Value: score.Value,
		}
	}
	for i, answerID := range answerIDs {
		if _, ok := scoreMap[answerID]; !ok {
			result[i] = &dataloader.Result[*model.Score]{Data: nil, Error: nil}
			continue
		}
		result[i] = &dataloader.Result[*model.Score]{Data: scoreMap[answerID], Error: nil}
	}
	return result
}

func LoadScoresByAnswerID(ctx context.Context, answerID string) (*model.Score, error) {
	loaders := GetLoaders(ctx)
	thunk := loaders.ScoreLoader.Load(ctx, answerID)
	return thunk()
}
