package loader

import "github.com/hytkgami/trivia-backend/usecase"

type Config struct {
	AnswerInteractor *usecase.AnswerInteractor
	UserInteractor   *usecase.UserInteractor
}
