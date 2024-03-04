package api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type APIFunc func(w http.ResponseWriter, r *http.Request)

type apiserver struct {
	listenAddress string
	router        chi.Router
	Handler       Handler
}

func New(listenAddress string) *apiserver {
	return &apiserver{
		listenAddress: listenAddress,
		router:        chi.NewRouter(),
	}
}

func (s *apiserver) Run() error {

	// sign up handler
	s.router.Post("/signup", handler(s.Handler.SignUp))

	// sign in handler
	s.router.Post("/signin", handler(s.Handler.SignIn))

	// start server & listen for incoming requests
	return http.ListenAndServe(s.listenAddress, s.router)
}

func handler(fn APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
		defer cancel()

		fn(w, r.WithContext(ctx))
	}
}
