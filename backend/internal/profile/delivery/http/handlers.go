package http

import (
	"encoding/json"
	"errors"
	"main/backend/internal/models"
	"main/backend/internal/profile"
	"main/backend/internal/session"
	"main/backend/internal/util"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type handlers struct {
	sessionRepository session.Repository
	profileRepository profile.Repository
}

func NewAuthHandlers(sR session.Repository, pR profile.Repository) *handlers {
	return &handlers{sR, pR}
}

func (h *handlers) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found, ok := r.Context().Value("user").(*models.User)

		if !ok {
			util.HandleError(w, errors.New("no user"))
			return
		}

		fields := &models.UserGeneralUpdate{}
		if err := json.NewDecoder(r.Body).Decode(fields); err != nil {
			util.HandleError(w, err)
			return
		}
		fields.Sanitize()
		if err := fields.Validate(); err != nil {
			util.HandleError(w, err)
			return
		}
		fields.UUID = found.UUID

		if fields.Avatar == found.Avatar && fields.DisplayName == found.Avatar && fields.About == found.Avatar {
			w.WriteHeader(http.StatusNotModified)
			return
		}

		if err := h.profileRepository.Update(r.Context(), fields); err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) PasswordUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found, ok := r.Context().Value("user").(*models.User)

		if !ok {
			util.HandleError(w, errors.New("no user"))
			return
		}

		fields := &models.UserPasswordUpdate{}
		if err := json.NewDecoder(r.Body).Decode(fields); err != nil {
			util.HandleError(w, err)
			return
		}
		if err := fields.Validate(); err != nil {
			util.HandleError(w, err)
			return
		}

		if err := fields.ComparePasswords(found.Password); err != nil {
			if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
				http.Error(w, "wrong credentials", http.StatusBadRequest)
				return
			}
			util.HandleError(w, err)
			return
		}
		if err := fields.HashPassword(); err != nil {
			util.HandleError(w, err)
			return
		}
		fields.UUID = found.UUID

		if err := h.profileRepository.PasswordUpdate(r.Context(), fields); err != nil {
			util.HandleError(w, err)
			return
		}

		session := &models.Session{UserID: found.UUID}
		if err := h.sessionRepository.InvalidateSessions(r.Context(), session); err != nil {
			util.HandleError(w, err)
			return
		}
		sessionId, err := h.sessionRepository.CreateSession(r.Context(), session, 30005)
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

func (h *handlers) DeleteAvatar() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found, ok := r.Context().Value("user").(*models.User)

		if !ok {
			util.HandleError(w, errors.New("no user"))
			return
		}

		if found.Avatar == "" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		avatar, err := h.profileRepository.DeleteAvatar(r.Context(), found)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		w.Write([]byte(avatar))
	}
}
