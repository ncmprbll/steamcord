package http

import (
	"encoding/json"
	"main/backend/internal/auth"
	"main/backend/internal/models"
	"main/backend/internal/session"
	"main/backend/pkg/http_errors"
	"net/http"
)

type handlers struct {
	authRepository auth.Repository
	sessionRepository session.Repository
}

func NewAuthHandlers(aR auth.Repository, sR session.Repository) *handlers {
	return &handlers{aR, sR}
}

func (h *handlers) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{}

		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			response := http_errors.ErrorResponse(err)
			http.Error(w, response.Error, response.Status)
			return
		}

		if err := user.HashPassword(); err != nil {
			response := http_errors.ErrorResponse(err)
			http.Error(w, response.Error, response.Status)
			return
		}

		if err := h.authRepository.Register(r.Context(), user); err != nil {
			response := http_errors.ErrorResponse(err)
			http.Error(w, response.Error, response.Status)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func (h *handlers) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{}

		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			response := http_errors.ErrorResponse(err)
			http.Error(w, response.Error, response.Status)
			return
		}

		found, err := h.authRepository.FindByLogin(r.Context(), user)

		if err != nil {
			response := http_errors.ErrorResponse(err)
			http.Error(w, response.Error, response.Status)
			return
		}

		if err := found.ComparePasswords(user.Password); err != nil {
			response := http_errors.ErrorResponse(err)
			http.Error(w, response.Error, response.Status)
			return
		}

		sessionId, err := h.sessionRepository.CreateSession(r.Context(), &models.Session{UserID: found.UUID}, 30)

		if err != nil {
			response := http_errors.ErrorResponse(err)
			http.Error(w, response.Error, response.Status)
			return
		}

		cookie := &http.Cookie{
			Name:     "session_id",
			Value:    sessionId,
			Path:     "/",
			MaxAge:   3600,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		}

		http.SetCookie(w, cookie)

		w.WriteHeader(http.StatusOK)
	}
}
