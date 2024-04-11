package http

import (
	"encoding/json"
	"main/backend/internal/models"
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
		var (
			rows []*models.GetTierRow
			err error
		)

		genre := r.URL.Query().Get("genre")
		count := r.URL.Query().Get("count")

		if count == "" {
			count = "14"
		}

		if genre == "" {
			rows, err = h.productsRepository.GetTier(r.Context(), count)
		} else {
			rows, err = h.productsRepository.GetTierByGenre(r.Context(), genre, count)
		}

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

func (h *handlers) GetFeatured() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := h.productsRepository.GetFeatured(r.Context())
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