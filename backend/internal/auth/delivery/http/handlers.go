package http

import (
	"encoding/json"
	"main/backend/internal/models"
	"main/backend/internal/auth"
	"net/http"
)

type handlers struct{
	repository auth.Repository
}

func NewAuthHandlers(repository auth.Repository) *handlers {
	return &handlers{repository: repository}
}

func (h *handlers) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{}

		// TODO: Remove error message with internal structure
		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// TODO: Remove error message with internal structure
		if err := h.repository.Register(r.Context(), user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}