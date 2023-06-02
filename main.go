package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/hytkgami/trivia-backend/graph"
	"github.com/hytkgami/trivia-backend/graph/loader"
	"github.com/hytkgami/trivia-backend/infrastructure"
	"github.com/hytkgami/trivia-backend/interfaces"
	"github.com/hytkgami/trivia-backend/interfaces/middleware"
	"github.com/hytkgami/trivia-backend/interfaces/repository"
	"github.com/hytkgami/trivia-backend/usecase"
	"github.com/rs/cors"
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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	})
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		AnswerInteractor: &usecase.AnswerInteractor{
			AnswerRepository: &repository.AnswerRepository{
				DB: db,
			},
		},
		UserInteractor: &usecase.UserInteractor{
			UserRepository: &repository.UserRepository{
				DB: db,
			},
		},
		LobbyInteractor: &usecase.LobbyInteractor{
			LobbyRepository: &repository.LobbyRepository{
				DB: db,
			},
		},
		QuestionInteractor: &usecase.QuestionInteractor{
			QuestionRepository: &repository.QuestionRepository{
				DB: db,
			},
		},
	}}))

	authHandler, err := infrastructure.NewFirebaseAuthHandler(ctx)
	if err != nil {
		return err
	}

	srv.AddTransport(&transport.POST{})
	srv.AddTransport(&transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			any := initPayload["Authorization"]
			authPayload, ok := any.(string)
			if !ok {
				return nil, errors.New("token not found in transport payaload")
			}
			idToken := strings.TrimPrefix(authPayload, "Bearer ")
			token, err := authHandler.VerifyIDToken(ctx, idToken)
			if err != nil {
				return nil, err
			}
			ctx = interfaces.SetUserUID(ctx, token.UID())
			return ctx, nil
		},
	})
	srv.Use(&extension.Introspection{})
	loaders := loader.NewLoaders(&loader.Config{
		AnswerInteractor: &usecase.AnswerInteractor{
			AnswerRepository: &repository.AnswerRepository{
				DB: db,
			},
		},
	})

	if os.Getenv("APP_ENV") == "development" {
		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	}
	http.Handle("/query", c.Handler(loader.Middleware(loaders, srv)))

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
