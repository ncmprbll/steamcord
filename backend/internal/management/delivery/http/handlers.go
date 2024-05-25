package http

import (
	"encoding/json"
	"main/backend/internal/management"
	"main/backend/internal/models"
	"main/backend/internal/util"
	"net/http"
)

type handlers struct {
	managementRepository management.Repository
}

func NewManagementHandlers(mR management.Repository) *handlers {
	return &handlers{mR}
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
		users, err := h.managementRepository.GetUsers(r.Context())
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
