package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/hytkgami/trivia-backend/domain"
	"github.com/hytkgami/trivia-backend/graph/loader"
	"github.com/hytkgami/trivia-backend/graph/model"
	"github.com/hytkgami/trivia-backend/interfaces"
	"github.com/hytkgami/trivia-backend/usecase"
)

// CreateQuestions is the resolver for the createQuestions field.
func (r *mutationResolver) CreateQuestions(ctx context.Context, lobbyID string, questions []*model.QuestionInput) (*model.CreateQuestionPayload, error) {
	uid, err := interfaces.GetUserUID(ctx)
	if err != nil {
		return nil, err
	}
	input := make([]*usecase.QuestionInput, len(questions))
	for i, q := range questions {
		input[i] = &usecase.QuestionInput{
			Title:       q.Title,
			OrderNumber: q.OrderNumber,
			Score:       q.Score,
		}
	}
	created, err := r.QuestionInteractor.CreateQuestions(ctx, uid, lobbyID, input)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Question, len(created))
	for i, q := range created {
		result[i] = &model.Question{
			ID:          q.ID,
			Title:       q.Title,
			OrderNumber: q.OrderNumber,
			Score:       q.Score,
		}
	}
	return &model.CreateQuestionPayload{
		Questions: result,
	}, nil
}

// PublishQuestion is the resolver for the publishQuestion field.
func (r *mutationResolver) PublishQuestion(ctx context.Context, lobbyID string, questionID string) (*model.PublishQuestionPayload, error) {
	q, err := r.QuestionInteractor.PublishQuestion(ctx, lobbyID, questionID)
	if err != nil {
		return nil, err
	}
	err = r.LobbyInteractor.PublishLobbyStatus(ctx, lobbyID, domain.LobbyStatusActive)
	if err != nil {
		return nil, err
	}
	return &model.PublishQuestionPayload{
		Question: &model.Question{
			ID:          q.ID,
			Title:       q.Title,
			OrderNumber: q.OrderNumber,
			Score:       q.Score,
		},
	}, nil
}

// Questions is the resolver for the questions field.
func (r *queryResolver) Questions(ctx context.Context, lobbyID string) ([]*model.Question, error) {
	questions, err := r.QuestionInteractor.FetchQuestionsByLobbyID(ctx, lobbyID)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Question, len(questions))
	for i, q := range questions {
		result[i] = &model.Question{
			ID:          q.ID,
			Title:       q.Title,
			OrderNumber: q.OrderNumber,
			Score:       q.Score,
		}
	}
	return result, nil
}

// Answers is the resolver for the answers field.
func (r *questionResolver) Answers(ctx context.Context, obj *model.Question) ([]*model.Answer, error) {
	return loader.LoadAnswersByQuestionID(ctx, obj.ID)
}

// CurrentQuestion is the resolver for the currentQuestion field.
func (r *subscriptionResolver) CurrentQuestion(ctx context.Context, lobbyID string) (<-chan *model.Question, error) {
	ch := make(chan *domain.Question)
	go func() {
		err := r.QuestionInteractor.SubscribeCurrentQuestion(ctx, lobbyID, ch)
		if err != nil {
			fmt.Println(err)
			return
		}
	}()
	modelCh := make(chan *model.Question)
	go func() {
		for q := range ch {
			m := &model.Question{
				ID:          q.ID,
				Title:       q.Title,
				OrderNumber: q.OrderNumber,
				Score:       q.Score,
			}
			select {
			case modelCh <- m:
			default:
				fmt.Println("Channel closed.")
				return
			}
		}
	}()
	return modelCh, nil
}

// Question returns QuestionResolver implementation.
func (r *Resolver) Question() QuestionResolver { return &questionResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type questionResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
