package http

import (
	"main/backend/internal/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(h *handlers, mw *middleware.MiddlewareManager) http.Handler {
	r := chi.NewRouter()

	r.Get("/{user_id}/comments", h.GetComments())

	r.Group(func(r chi.Router) {
		r.Use(mw.AuthSessionMiddleware)
		r.Post("/{user_id}/friend-invite", h.FriendInvite())
		r.Post("/{user_id}/friend-reject", h.HandleFriendInvite("rejected"))
		r.Post("/{user_id}/friend-accept", h.HandleFriendInvite("accepted"))
		r.Post("/{user_id}/comments", h.CreateComment())
		r.Patch("/", h.Update())
		r.Patch("/password", h.PasswordUpdate())
		r.Patch("/privacy", h.PrivacyUpdate())
		r.Delete("/avatar", h.DeleteAvatar())
	})

	return r
}
