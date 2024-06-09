package middleware

import (
	"context"
	"main/backend/internal/models"
	"net/http"
)

func (mw *MiddlewareManager) GetUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionIdCookie, err := r.Cookie("session_id")
		if err != nil {
			next.ServeHTTP(w, r.WithContext(r.Context()))
			return
		}

		sessionId := sessionIdCookie.Value
		session, err := mw.sessionRepository.GetSessionByID(r.Context(), sessionId)
		if err != nil {
			next.ServeHTTP(w, r.WithContext(r.Context()))
			return
		}

		found, err := mw.authRepository.FindByUUID(r.Context(), &models.User{UUID: session.UserID})
		if err != nil {
			next.ServeHTTP(w, r.WithContext(r.Context()))
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user", found)))
	})
}

func (mw *MiddlewareManager) AuthSessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionIdCookie, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		sessionId := sessionIdCookie.Value
		session, err := mw.sessionRepository.GetSessionByID(r.Context(), sessionId)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		found, err := mw.authRepository.FindByUUID(r.Context(), &models.User{UUID: session.UserID})
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user", found)))
	})
}
