package middleware

import (
	"context"
	"main/backend/internal/models"
	"main/backend/internal/util"
	"net/http"
)


func (mw *MiddlewareManager) AuthSessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionIdCookie, err := r.Cookie("session_id")

		if err != nil {
			util.HandleError(w, err)
			return
		}

		sessionId := sessionIdCookie.Value
		session, err := mw.sessionRepository.GetSessionByID(r.Context(), sessionId)

		if err != nil {
			util.HandleError(w, err)
			return
		}

		found, err := mw.authRepository.FindByUUID(r.Context(), &models.User{UUID: session.UserID})
		if err != nil {
			util.HandleError(w, err)
			return
		}
		found.SanitizePassword()

		ctx := context.WithValue(r.Context(), "user", found)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}