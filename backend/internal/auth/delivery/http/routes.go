package http

import (
	"main/backend/internal/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(h *handlers, mw *middleware.MiddlewareManager) http.Handler {
	r := chi.NewRouter()

	r.Post("/register", h.Register())
	r.Post("/login", h.Login())
	r.Get("/{user_id}", h.FindByUUID())

	r.Group(func(r chi.Router) {
		r.Use(mw.AuthSessionMiddleware)
		r.Get("/me", h.Me())
	})

	return r
}
