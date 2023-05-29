package loader

import (
	"context"

	"github.com/graph-gophers/dataloader/v7"
	"github.com/hytkgami/trivia-backend/graph/model"
	"github.com/hytkgami/trivia-backend/usecase"
)

type AnswerLoader struct {
	Interactor *usecase.AnswerInteractor
}

func (l *AnswerLoader) BatchGetByQuestionIDs(ctx context.Context, questionIDs []string) []*dataloader.Result[[]*model.Answer] {
	result := make([]*dataloader.Result[[]*model.Answer], len(questionIDs))
	answers, err := l.Interactor.FetchByQuestionIDs(ctx, questionIDs)
	if err != nil {
		for i := range questionIDs {
			result[i] = &dataloader.Result[[]*model.Answer]{Data: nil, Error: err}
		}
		return result
	}
	mapped := map[string][]*model.Answer{}
	for _, answer := range answers {
		ans := &model.Answer{
			ID:      answer.ID,
			Content: answer.Content,
		}
		mapped[answer.QuestionID] = append(mapped[answer.QuestionID], ans)
	}
	for i, questionID := range questionIDs {
		answers, ok := mapped[questionID]
		if !ok {
			answers = []*model.Answer{}
		}
		result[i] = &dataloader.Result[[]*model.Answer]{Data: answers, Error: nil}
	}
	return result
}

func LoadAnswersByQuestionID(ctx context.Context, questionID string) ([]*model.Answer, error) {
	loaders := GetLoaders(ctx)
	thunk := loaders.AnswerLoader.Load(ctx, questionID)
	return thunk()
}
