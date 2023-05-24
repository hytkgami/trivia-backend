package middleware_test

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hytkgami/trivia-backend/interfaces/middleware"
)

func TestValidate(t *testing.T) {
	type testCase struct {
		name        string
		headerKey   string
		headerValue string
		want        bool
	}
	testCases := []testCase{
		{
			name:        "valid header",
			headerKey:   "Authorization",
			headerValue: "Bearer eyJhbGciOiJSUzI1NiIsIm",
			want:        true,
		},
		{
			name:        "invalid header",
			headerKey:   "Authorization",
			headerValue: "invalid",
			want:        false,
		},
	}
	mw := &middleware.AuthMiddleware{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/", nil)
			if err != nil {
				t.Errorf("failed to create request: %v", err)
			}
			req.Header.Add(tc.headerKey, tc.headerValue)
			got := middleware.Validate(mw, req)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Validate() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestIsPublicPath(t *testing.T) {
	type testCase struct {
		path string
		want bool
	}
	testCases := []testCase{
		{
			path: "/ping",
			want: true,
		},
		{
			path: "/",
			want: true,
		},
		{
			path: "/query",
			want: false,
		},
	}
	mw := &middleware.AuthMiddleware{}
	for _, tc := range testCases {
		t.Run(tc.path, func(t *testing.T) {
			got := middleware.IsPublicPath(mw, tc.path)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("IsPublicPath() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
