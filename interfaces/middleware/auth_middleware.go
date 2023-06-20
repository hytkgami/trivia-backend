package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/hytkgami/trivia-backend/interfaces"
	"github.com/hytkgami/trivia-backend/interfaces/repository"
)

type AuthMiddleware struct {
	authHandler repository.FirebaseAuthHandler
}

func NewAuthMiddleware(authHandler repository.FirebaseAuthHandler) *AuthMiddleware {
	return &AuthMiddleware{
		authHandler: authHandler,
	}
}

func (m *AuthMiddleware) isWebSocket(r *http.Request) bool {
	return r.Header.Get("Upgrade") == "websocket" && r.Header.Get("Connection") == "Upgrade"
}

func (m *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("headers", r.Header)
		if m.isPublicPath(r.URL.Path) {
			log.Println("public path")
			next.ServeHTTP(w, r)
			return
		}
		if r.Method == http.MethodOptions {
			log.Println("options method")
			next.ServeHTTP(w, r)
			return
		}
		if m.isWebSocket(r) {
			log.Println("websocket")
			next.ServeHTTP(w, r)
			return
		}
		if !m.validate(r) {
			log.Println("invalid headers")
			interfaces.HttpErrorResponse(w, fmt.Errorf("invalid headers"), http.StatusBadRequest)
			return
		}
		idToken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		token, err := m.authHandler.VerifyIDToken(r.Context(), idToken)
		if err != nil {
			log.Println("unauthorized")
			interfaces.HttpErrorResponse(w, fmt.Errorf("unauthorized: %v", err), http.StatusUnauthorized)
			return
		}
		ctx := interfaces.SetUserUID(r.Context(), token.UID())
		log.Println("authorized")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (*AuthMiddleware) isPublicPath(path string) bool {
	paths := map[string]struct{}{
		"/ping": {},
		"/":     {},
	}
	_, ok := paths[path]
	return ok
}

func (*AuthMiddleware) validate(r *http.Request) bool {
	return strings.HasPrefix(r.Header.Get("Authorization"), "Bearer ")
}
