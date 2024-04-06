package http

import (
	"encoding/json"
	"main/backend/internal/auth"
	"main/backend/internal/models"
	"main/backend/pkg/http_errors"
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

		if err := h.repository.Register(r.Context(), user); err != nil {
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

		found, err := h.repository.FindByLogin(r.Context(), user)

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

		// Create session

		w.WriteHeader(http.StatusOK)
	}
}