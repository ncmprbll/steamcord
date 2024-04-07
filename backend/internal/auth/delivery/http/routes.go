package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(h *handlers) http.Handler {
	r := chi.NewRouter()

	r.Post("/register", h.Register())
	r.Post("/login", h.Login())
	r.Get("/{user_id}", h.FindByUUID())
	r.Get("/me", h.Me())

	return r
}
