package graph

import (
	"github.com/hytkgami/trivia-backend/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AnswerInteractor   *usecase.AnswerInteractor
	UserInteractor     *usecase.UserInteractor
	LobbyInteractor    *usecase.LobbyInteractor
	QuestionInteractor *usecase.QuestionInteractor
	ScoreInteractor    *usecase.ScoreInteractor
}
