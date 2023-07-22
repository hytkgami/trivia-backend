package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"
	"strings"

	"github.com/hytkgami/trivia-backend/graph/loader"
	"github.com/hytkgami/trivia-backend/graph/model"
	"github.com/hytkgami/trivia-backend/interfaces"
)

// Owner is the resolver for the owner field.
func (r *answerResolver) Owner(ctx context.Context, obj *model.Answer) (*model.User, error) {
	return loader.LoadUser(ctx, obj.UID)
}

// Score is the resolver for the score field.
func (r *answerResolver) Score(ctx context.Context, obj *model.Answer) (*model.Score, error) {
	return loader.LoadScoresByAnswerID(ctx, obj.ID)
}

// Answer is the resolver for the answer field.
func (r *mutationResolver) Answer(ctx context.Context, questionID string, answer string) (*model.AnswerPayload, error) {
	uid, err := interfaces.GetUserUID(ctx)
	if err != nil {
		return nil, err
	}
	a, err := r.AnswerInteractor.Create(ctx, uid, questionID, answer)
	if err != nil {
		return nil, err
	}
	return &model.AnswerPayload{
		Answer: &model.Answer{
			ID:         a.ID,
			Content:    a.Content,
			QuestionID: a.QuestionID,
			UID:        a.UID,
		},
	}, nil
}

// Scoring is the resolver for the scoring field.
func (r *mutationResolver) Scoring(ctx context.Context, answerID string, mark model.Mark, value int) (*model.ScoringPayload, error) {
	markString := strings.ToLower(string(mark))
	err := r.ScoreInteractor.CreateScore(ctx, answerID, markString, value)
	if err != nil {
		return nil, err
	}
	answer, err := r.AnswerInteractor.FetchByID(ctx, answerID)
	if err != nil {
		return nil, err
	}
	return &model.ScoringPayload{
		Answer: &model.Answer{
			ID:         answer.ID,
			Content:    answer.Content,
			QuestionID: answer.QuestionID,
			UID:        answer.UID,
		},
	}, nil
}

// Answer returns AnswerResolver implementation.
func (r *Resolver) Answer() AnswerResolver { return &answerResolver{r} }

type answerResolver struct{ *Resolver }
