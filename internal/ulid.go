package internal

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
)

func NewULID() (string, error) {
	entropy := ulid.Monotonic(rand.Reader, 0)
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, entropy)
	if err != nil {
		return "", fmt.Errorf("failed to generate ulid: %w", err)
	}
	return id.String(), nil
}
