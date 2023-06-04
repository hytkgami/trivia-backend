package usecase

import (
	"context"
	"fmt"

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

func (i *QuestionInteractor) PublishQuestion(ctx context.Context, lobbyID, questionID string) (*domain.Question, error) {
	err := i.QuestionRepository.PublishQuestion(ctx, lobbyID, questionID)
	if err != nil {
		return nil, err
	}
	return i.QuestionRepository.FetchQuestionByID(ctx, questionID)
}

func (i *QuestionInteractor) FetchQuestionByID(ctx context.Context, questionID string) (*domain.Question, error) {
	return i.QuestionRepository.FetchQuestionByID(ctx, questionID)
}

func (i *QuestionInteractor) FetchQuestionsByLobbyID(ctx context.Context, lobbyID string) ([]*domain.Question, error) {
	return i.QuestionRepository.FetchQuestionsByLobbyID(ctx, lobbyID)
}

func (i *QuestionInteractor) SubscribeCurrentQuestion(ctx context.Context, lobbyID string, ch chan<- *domain.Question) error {
	questionID, err := i.QuestionRepository.FetchCurrentQuestionID(ctx, lobbyID)
	if err != nil {
		return err
	}
	question, err := i.QuestionRepository.FetchQuestionByID(ctx, questionID)
	if err != nil {
		return err
	}
	ch <- question
	questionIDCh := make(chan string)
	i.QuestionRepository.SubscribeCurrentQuestionID(ctx, lobbyID, questionIDCh)
	go func() {
		for questionID := range questionIDCh {
			question, err := i.QuestionRepository.FetchQuestionByID(ctx, questionID)
			if err != nil {
				continue
			}
			select {
			case ch <- question:
			default:
				fmt.Println("question channel is full")
				return
			}
		}
	}()
	return nil
}
