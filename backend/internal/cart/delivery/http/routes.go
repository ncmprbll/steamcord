package http

import (
	"main/backend/internal/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(h *handlers, mw *middleware.MiddlewareManager) http.Handler {
	r := chi.NewRouter()

	r.Use(mw.AuthSessionMiddleware)

	r.Post("/", h.AddToCart())
	r.Post("/purchase", h.Purchase())
	r.Get("/ids", h.CartIDs())
	r.Get("/", h.Cart())
	r.Delete("/", h.DeleteFromCart())

	return r
}
