package http

import (
	"main/backend/internal/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(h *handlers, mw *middleware.MiddlewareManager) http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(mw.AuthSessionMiddleware)
		r.Patch("/", h.Update())
		r.Patch("/password", h.PasswordUpdate())
	})

	return r
}
