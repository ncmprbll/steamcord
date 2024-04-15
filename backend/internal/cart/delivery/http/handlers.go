package http

import (
	"encoding/json"
	"errors"
	"main/backend/internal/auth"
	"main/backend/internal/cart"
	"main/backend/internal/models"
	"main/backend/internal/session"
	"main/backend/internal/util"
	"net/http"
)

type handlers struct {
	cartRepository cart.Repository
	authRepository auth.Repository
	sessionRepository session.Repository
}

func NewAuthHandlers(cR cart.Repository, aR auth.Repository, sR session.Repository) *handlers {
	return &handlers{cR, aR, sR}
}

func (h *handlers) Cart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found, ok := r.Context().Value("user").(*models.User)

		if !ok {
			util.HandleError(w, errors.New("no user"))
			return
		}

		rows, err := h.cartRepository.Cart(r.Context(), found)
		if err != nil {
			util.HandleError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(rows); err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) CartCount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found, ok := r.Context().Value("user").(*models.User)

		if !ok {
			util.HandleError(w, errors.New("no user"))
			return
		}

		cartJson, err := h.cartRepository.CartCount(r.Context(), found)
		if err != nil {
			util.HandleError(w, err)
			return	
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(cartJson); err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h *handlers) AddToCart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found, ok := r.Context().Value("user").(*models.User)

		if !ok {
			util.HandleError(w, errors.New("no user"))
			return
		}

		jsonProductID := &models.JSONProductID{}

		if err := json.NewDecoder(r.Body).Decode(jsonProductID); err != nil {
			util.HandleError(w, err)
			return
		}

		if err := h.cartRepository.AddToCart(r.Context(), &models.CartRow{UserID: found.UUID, ProductID: jsonProductID.ProductID}); err != nil {
			util.HandleError(w, err)
			return	
		}

		w.WriteHeader(http.StatusOK)
	}
}