package repository

import (
	"context"
)

type FirebaseAuthHandler interface {
	VerifyIDToken(ctx context.Context, idToken string) (FirebaseAuthToken, error)
}

type FirebaseAuthToken interface {
	UID() string
}
