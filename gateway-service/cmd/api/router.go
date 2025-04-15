package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/microservice-monorepo/gateway-service/internal/gateway/interfaces"
)

type system struct {
	Config  *config
	Handler *handlers
}

type config struct {
	Addr string
}

type handlers struct {
	Signup interfaces.SignupHandler
}

func (app system) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Mount some handlers
	r.Mount("/v1", interfaces.NewGatewayRouter())

	return r
}

func (app system) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      mux,
		IdleTimeout:  time.Second * 30,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	return srv.ListenAndServe()
}
