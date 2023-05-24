package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hytkgami/trivia-backend/graph"
	"github.com/hytkgami/trivia-backend/infrastructure"
	"github.com/hytkgami/trivia-backend/interfaces/middleware"
	"github.com/hytkgami/trivia-backend/interfaces/repository"
	"github.com/hytkgami/trivia-backend/usecase"
	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatalf("failed to run: %+v", err)
	}
}

const defaultPort = "8080"

func run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}))

	db, err := infrastructure.NewDB(ctx)
	if err != nil {
		return err
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		UserInteractor: &usecase.UserInteractor{
			UserRepository: &repository.UserRepository{
				DB: db,
			},
		},
	}}))

	if os.Getenv("APP_ENV") == "development" {
		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	}
	http.Handle("/query", srv)

	authHandler, err := infrastructure.NewFirebaseAuthHandler(ctx)
	if err != nil {
		return err
	}
	authMiddleware := middleware.NewAuthMiddleware(authHandler)
	s := &http.Server{
		Addr:    ":" + port,
		Handler: authMiddleware.Middleware(http.DefaultServeMux),
	}

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		log.Println("Listening on port", port)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})

	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}

	return eg.Wait()
}
