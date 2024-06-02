package http

import (
	"encoding/json"
	"main/backend/internal/language"
	"main/backend/internal/util"
	"net/http"
)

type handlers struct {
	languageRepository language.Repository
}

func NewAuthHandlers(lR language.Repository) *handlers {
	return &handlers{lR}
}

func (h *handlers) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		locales, err := h.languageRepository.GetAll(r.Context())
		if err != nil {
			util.HandleError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(locales)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
