package http

import (
	"main/backend/internal/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(h *handlers, mw *middleware.MiddlewareManager) http.Handler {
	r := chi.NewRouter()

	r.Use(mw.GetUserMiddleware)

	r.Get("/tier", h.GetTier())
	r.Get("/featured", h.GetFeatured())

	return r
}
