package http

import (
	"main/backend/internal/middleware"
	"main/backend/internal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(h *handlers, mw *middleware.MiddlewareManager) http.Handler {
	r := chi.NewRouter()

	r.Use(mw.AuthSessionMiddleware)

	r.Group(func(r chi.Router) {
		r.Get("/permissions", h.GetPermissions())
	})

	r.Group(func(r chi.Router) {
		r.Use(mw.HasPermissionsMiddleware(&models.Permissions{models.PERMISSION_USERS_MANAGEMENT}))
		r.Get("/users", h.GetUsers())
		r.Patch("/users/{user_id}", h.UpdateUser())
	})

	r.Group(func(r chi.Router) {
		r.Use(mw.HasPermissionsMiddleware(&models.Permissions{models.PERMISSION_ROLES_MANAGEMENT}))
		r.Get("/roles", h.GetRoles())
		r.Get("/roles/permissions", h.GetRolePermissions())
		r.Post("/roles", h.CreateRole())
		r.Delete("/roles/{role_id}", h.DeleteRole())
	})

	return r
}
