package http

import (
	"encoding/json"
	"main/backend/internal/auth"
	"main/backend/internal/models"
	"main/backend/internal/profile"
	"main/backend/internal/session"
	"main/backend/internal/util"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type handlers struct {
	authRepository    auth.Repository
	sessionRepository session.Repository
	profileRepository profile.Repository
}

func NewAuthHandlers(aR auth.Repository, sR session.Repository, pR profile.Repository) *handlers {
	return &handlers{aR, sR, pR}
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
		lastIndex := strings.LastIndex(r.RemoteAddr, ":")
		ip := r.Header.Get("X-Real-IP")
		if ip == "" {
			ip = r.RemoteAddr[:lastIndex]
		}

		attempts, err := h.sessionRepository.GetIPBadLoginAttempts(r.Context(), ip)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		if attempts >= models.MAX_BAD_LOGIN_ATTEMPTS {
			http.Error(w, "too many attempts, timeout", http.StatusBadRequest)
			return
		}

		user := &models.User{}
		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			util.HandleError(w, err)
			return
		}

		found, err := h.authRepository.FindByLogin(r.Context(), user)
		if err != nil {
			h.sessionRepository.LogIPBadLoginAttempt(r.Context(), ip, models.BAD_LOGIN_IP_TIMEOUT)
			http.Error(w, "wrong credentials", http.StatusNotFound)
			return
		}

		if err := found.ComparePasswords(user.Password); err != nil {
			h.sessionRepository.LogIPBadLoginAttempt(r.Context(), ip, models.BAD_LOGIN_IP_TIMEOUT)
			http.Error(w, "wrong credentials", http.StatusNotFound)
			return
		}
		found.SanitizePassword()

		if *found.Banned {
			http.Error(w, "banned", http.StatusConflict)
			return
		}

		if err := h.authRepository.UpdateLoginDate(r.Context(), found); err != nil {
			util.HandleError(w, err)
			return
		}

		sessionId, err := h.sessionRepository.CreateSession(r.Context(), &models.Session{UserID: found.UUID}, 24*60*60)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		cookie := &http.Cookie{
			Name:     "session_id",
			Value:    sessionId,
			Path:     "/",
			MaxAge:   86400,
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
		}
		http.SetCookie(w, cookie)
		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) FindByUUID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requester, ok := r.Context().Value("user").(*models.User)
		userId := chi.URLParam(r, "user_id")
		uuid, err := uuid.Parse(userId)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		user := &models.User{UUID: uuid}
		found, err := h.authRepository.FindByUUID(r.Context(), user)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		if found.Privacy == "private" {
			if !ok || requester.UUID != found.UUID {
				found.ApplyPrivacy()
			}
		} else if found.Privacy == "friendsOnly" {
			if ok && requester.UUID != found.UUID {
				friends, err := h.profileRepository.IsFriend(r.Context(), user, requester)
				if err != nil {
					util.HandleError(w, err)
					return
				}
				if !friends {
					found.ApplyPrivacy()
				}
			} else if !ok {
				found.ApplyPrivacy()
			}
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
		found := r.Context().Value("user").(*models.User)
		found.SanitizePassword()

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
		found := r.Context().Value("user").(*models.User)
		sessionIdCookie, err := r.Cookie("session_id")

		if err != nil {
			util.HandleError(w, err)
			return
		}

		if err := h.sessionRepository.DeleteSession(r.Context(), &models.Session{SessionID: sessionIdCookie.Value, UserID: found.UUID}); err != nil {
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
