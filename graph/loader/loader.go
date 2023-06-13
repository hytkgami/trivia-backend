package loader

import (
	"context"
	"net/http"

	"github.com/graph-gophers/dataloader/v7"
	"github.com/hytkgami/trivia-backend/graph/model"
)

type loaderKey struct{}

type Loaders struct {
	AnswerLoader *dataloader.Loader[string, []*model.Answer]
	ScoreLoader  *dataloader.Loader[string, *model.Score]
	UserLoader   *dataloader.Loader[string, *model.User]
}

func Middleware(loaders *Loaders, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), loaderKey{}, loaders)
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}

func NewLoaders(cfg *Config) *Loaders {
	answerLoader := &AnswerLoader{
		Interactor: cfg.AnswerInteractor,
	}
	scoreLoader := &ScoreLoader{
		ScoreInteractor: cfg.ScoreInteractor,
	}
	userLoader := &UserLoader{
		Interactor: cfg.UserInteractor,
	}
	return &Loaders{
		AnswerLoader: dataloader.NewBatchedLoader(
			answerLoader.BatchGetByQuestionIDs,
		),
		ScoreLoader: dataloader.NewBatchedLoader(
			scoreLoader.BatchGetByAnswerIDs,
		),
		UserLoader: dataloader.NewBatchedLoader(
			userLoader.BatchGet,
		),
	}
}

func GetLoaders(ctx context.Context) *Loaders {
	return ctx.Value(loaderKey{}).(*Loaders)
}
