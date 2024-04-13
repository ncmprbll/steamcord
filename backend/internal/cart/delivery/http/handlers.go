package http

import (
	"encoding/json"
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

func (h *handlers) AddToCart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionIdCookie, err := r.Cookie("session_id")

		if err != nil {
			util.HandleError(w, err)
			return
		}

		sessionId := sessionIdCookie.Value
		session, err := h.sessionRepository.GetSessionByID(r.Context(), sessionId)

		if err != nil {
			util.HandleError(w, err)
			return
		}

		found, err := h.authRepository.FindByUUID(r.Context(), &models.User{UUID: session.UserID})
		if err != nil {
			util.HandleError(w, err)
			return
		}
		found.SanitizePassword()

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