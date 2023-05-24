package interfaces

import (
	"context"
	"errors"
)

type userUIDKey struct{}

func SetUserUID(ctx context.Context, uid string) context.Context {
	return context.WithValue(ctx, userUIDKey{}, uid)
}

func GetUserUID(ctx context.Context) (string, error) {
	uid, ok := ctx.Value(userUIDKey{}).(string)
	if !ok {
		return "", errors.New("user uid not found in context")
	}
	return uid, nil
}
