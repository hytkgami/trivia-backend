package infrastructure

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/hytkgami/trivia-backend/interfaces/repository"
)

type FirebaseAuthHandler struct {
	client *auth.Client
}

func NewFirebaseAuthHandler(ctx context.Context) (*FirebaseAuthHandler, error) {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, err
	}
	auth, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	return &FirebaseAuthHandler{client: auth}, nil
}

type FirebaseAuthToken struct {
	token *auth.Token
}

func (h *FirebaseAuthHandler) VerifyIDToken(ctx context.Context, idToken string) (repository.FirebaseAuthToken, error) {
	token, err := h.client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}
	return &FirebaseAuthToken{token: token}, nil
}

func (t *FirebaseAuthToken) UID() string {
	return t.token.UID
}
