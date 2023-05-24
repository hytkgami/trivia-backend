package internal

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func NewULID() (string, error) {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, entropy)
	if err != nil {
		return "", fmt.Errorf("failed to generate ulid: %w", err)
	}
	return id.String(), nil
}
