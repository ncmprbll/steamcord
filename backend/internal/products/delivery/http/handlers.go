package http

import (
	"encoding/json"
	"errors"
	"main/backend/internal/models"
	"main/backend/internal/products"
	"main/backend/internal/util"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
			currencyCode = "USD"
			rows         []*models.TierRow
			err          error
		)

		found, ok := r.Context().Value("user").(*models.User)

		if ok {
			currencyCode = found.CurrencyCode
		}

		genre := r.URL.Query().Get("genre")
		count := r.URL.Query().Get("count")

		if count == "" {
			count = "14"
		}

		if genre == "" {
			rows, err = h.productsRepository.GetTier(r.Context(), currencyCode, count)
		} else {
			rows, err = h.productsRepository.GetTierByGenre(r.Context(), currencyCode, genre, count)
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
		currencyCode := "USD"
		found, ok := r.Context().Value("user").(*models.User)

		if ok {
			currencyCode = found.CurrencyCode
		}

		rows, err := h.productsRepository.GetFeatured(r.Context(), currencyCode)
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

func (h *handlers) GetOwned() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found, ok := r.Context().Value("user").(*models.User)

		if !ok {
			util.HandleError(w, errors.New("no user"))
			return
		}

		ownedJson, err := h.productsRepository.GetOwnedIDs(r.Context(), found)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(ownedJson); err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) FindByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productId := chi.URLParam(r, "product_id")

		i, err := strconv.Atoi(productId)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		currencyCode := "USD"
		found, ok := r.Context().Value("user").(*models.User)

		if ok {
			currencyCode = found.CurrencyCode
		}

		product, err := h.productsRepository.FindByID(r.Context(), &models.Product{ID: i}, currencyCode)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(product); err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
