package http

import (
	"encoding/json"
	"errors"
	"main/backend/internal/auth"
	"main/backend/internal/models"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
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
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == "23505" {
				http.Error(w, err.Error(), http.StatusConflict)
				return
			}

			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}