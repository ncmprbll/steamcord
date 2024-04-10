package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(h *handlers) http.Handler {
	r := chi.NewRouter()

	r.Get("/tier", h.GetTier())
	r.Get("/featured", h.GetFeatured())

	return r
}
