package http

import (
	"encoding/json"
	"main/backend/internal/products"
	"main/backend/internal/util"
	"net/http"
)

type handlers struct {
	productsRepository products.Repository
}

func NewAuthHandlers(pR products.Repository) *handlers {
	return &handlers{pR}
}

func (h *handlers) GetTier() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := h.productsRepository.GetTier(r.Context(), 14)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(rows)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}