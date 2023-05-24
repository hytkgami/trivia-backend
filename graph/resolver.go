package graph

import "github.com/hytkgami/trivia-backend/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserInteractor  *usecase.UserInteractor
	LobbyInteractor *usecase.LobbyInteractor
}
