package http

import (
	"encoding/json"
	"errors"
	"main/backend/internal/auth"
	"main/backend/internal/models"
	"main/backend/internal/session"
	"main/backend/internal/util"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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
			util.HandleError(w, err)
			return
		}

		if err := user.Validate(); err != nil {
			util.HandleError(w, err)
			return
		}

		if err := user.HashPassword(); err != nil {
			util.HandleError(w, err)
			return
		}

		if err := h.authRepository.Register(r.Context(), user); err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func (h *handlers) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{}

		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			util.HandleError(w, err)
			return
		}

		found, err := h.authRepository.FindByLogin(r.Context(), user)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		if err := found.ComparePasswords(user.Password); err != nil {
			util.HandleError(w, err)
			return
		}
		found.SanitizePassword()

		sessionId, err := h.sessionRepository.CreateSession(r.Context(), &models.Session{UserID: found.UUID}, 30005)

		if err != nil {
			util.HandleError(w, err)
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

func (h *handlers) FindByUUID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := chi.URLParam(r, "user_id")

		uuid, err := uuid.Parse(userId)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		found, err := h.authRepository.FindByUUID(r.Context(), &models.User{UUID: uuid})
		if err != nil {
			util.HandleError(w, err)
			return
		}
		found.RemoveSensitiveData()

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(found)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) Me() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found, ok := r.Context().Value("user").(*models.User)

		if !ok {
			util.HandleError(w, errors.New("no user"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(found); err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionIdCookie, err := r.Cookie("session_id")

		if err != nil {
			util.HandleError(w, err)
			return
		}

		if err := h.sessionRepository.DeleteByID(r.Context(), sessionIdCookie.Value); err != nil {
			util.HandleError(w, err)
			return
		}

		cookie := &http.Cookie{
			Name:     "session_id",
			Value:    "",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
		}

		http.SetCookie(w, cookie)
		w.WriteHeader(http.StatusOK)
	}
}

