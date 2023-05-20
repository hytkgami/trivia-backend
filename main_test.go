package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestHealthCheck(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx)
	})

	port := "10080"
	t.Setenv("PORT", port)

	resp, err := http.Get(fmt.Sprintf("http://localhost:%s/ping", port))
	if err != nil {
		t.Errorf("failed to health check: %+v", err)
	}
	defer resp.Body.Close()

	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	want := "pong"
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}

	cancel()
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
