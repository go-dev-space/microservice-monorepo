package interfaces

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewGatewayRouter() http.Handler {
	r := chi.NewRouter()

	// auth handlers
	r.Route("/auth", func(r chi.Router) {

	})

	return r
}
