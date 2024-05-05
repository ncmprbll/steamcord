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

		if err := h.sessionRepository.InvalidateSessions(r.Context(), &models.Session{UserID: found.UUID}); err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
