package http

import (
	"encoding/json"
	"main/backend/internal/management"
	"main/backend/internal/models"
	"main/backend/internal/session"
	"main/backend/internal/util"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type handlers struct {
	managementRepository management.Repository
	sessionRepository    session.Repository
}

func NewManagementHandlers(mR management.Repository, sR session.Repository) *handlers {
	return &handlers{mR, sR}
}

func (h *handlers) GetPermissions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found := r.Context().Value("user").(*models.User)

		permissions, err := h.managementRepository.GetPermissions(r.Context(), found)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(permissions); err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) GetUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		term := r.URL.Query().Get("term")
		users, err := h.managementRepository.GetUsers(r.Context(), term)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(users); err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{}
		userId := chi.URLParam(r, "user_id")
		uuid, err := uuid.Parse(userId)
		if err != nil {
			util.HandleError(w, err)
			return
		}
		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			util.HandleError(w, err)
			return
		}
		user.UUID = uuid

		if user.Password != "" {
			if err := user.HashPassword(); err != nil {
				util.HandleError(w, err)
				return
			}
		}

		if err := h.managementRepository.UpdateUser(r.Context(), user); err != nil {
			util.HandleError(w, err)
			return
		}

		if user.Banned != nil && *user.Banned || user.Password != "" {
			if err := h.sessionRepository.InvalidateSessions(r.Context(), &models.Session{UserID: user.UUID}); err != nil {
				util.HandleError(w, err)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) GetRoles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roles, err := h.managementRepository.GetRoles(r.Context())
		if err != nil {
			util.HandleError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(roles); err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) CreateRole() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role := &models.Role{}
		if err := json.NewDecoder(r.Body).Decode(role); err != nil {
			util.HandleError(w, err)
			return
		}

		if err := role.Validate(); err != nil {
			if strings.Contains(err.Error(), "illegal") {
				http.Error(w, "bad name", http.StatusConflict)
			} else {
				util.HandleError(w, err)
			}
			return
		}

		if err := h.managementRepository.CreateRole(r.Context(), role); err != nil {
			if strings.Contains(err.Error(), "23505") {
				http.Error(w, "role exists", http.StatusConflict)
			} else {
				util.HandleError(w, err)
			}
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) DeleteRole() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roleId := chi.URLParam(r, "role_id")
		id, err := strconv.Atoi(roleId)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		affected, err := h.managementRepository.DeleteRole(r.Context(), &models.Role{ID: id})
		if err != nil {
			util.HandleError(w, err)
			return
		}

		if affected == 0 {
			http.Error(w, "bad role", http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

func (h *handlers) GetRolePermissions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := h.managementRepository.GetRolePermissions(r.Context())
		if err != nil {
			util.HandleError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) AddPermission() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roleId := chi.URLParam(r, "role_id")
		id, err := strconv.Atoi(roleId)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		permissions := &models.Permissions{}
		if err := json.NewDecoder(r.Body).Decode(permissions); err != nil {
			util.HandleError(w, err)
			return
		}

		if err := h.managementRepository.AddPermission(r.Context(), &models.Role{ID: id}, permissions); err != nil {
			if strings.Contains(err.Error(), "23503") {
				http.Error(w, "bad permission", http.StatusBadRequest)
			} else {
				util.HandleError(w, err)
			}
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) DeletePermission() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roleId := chi.URLParam(r, "role_id")
		id, err := strconv.Atoi(roleId)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		permissions := &models.Permissions{}
		if err := json.NewDecoder(r.Body).Decode(permissions); err != nil {
			util.HandleError(w, err)
			return
		}

		affected, err := h.managementRepository.DeletePermission(r.Context(), &models.Role{ID: id}, permissions)
		if err != nil {
			if strings.Contains(err.Error(), "23503") {
				http.Error(w, "bad permission", http.StatusBadRequest)
			} else {
				util.HandleError(w, err)
			}
			return
		}

		if affected == 0 {
			http.Error(w, "no permissions found", http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}
