package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"main/backend/internal/models"
	"main/backend/internal/profile"
	"main/backend/internal/util"
	"net/http"
)

type handlers struct {
	profileRepository profile.Repository
}

func NewAuthHandlers(pR profile.Repository) *handlers {
	return &handlers{pR}
}

func (h *handlers) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found, ok := r.Context().Value("user").(*models.User)

		if !ok {
			util.HandleError(w, errors.New("no user"))
			return
		}

		fields := &models.UserGeneralUpdate{}
		if err := json.NewDecoder(r.Body).Decode(fields); err != nil {
			util.HandleError(w, err)
			return
		}
		fields.Sanitize()
		if err := fields.Validate(); err != nil {
			util.HandleError(w, err)
			return
		}
		fields.UUID = found.UUID

		fmt.Println(fields)

		if fields.Avatar == "" && fields.DisplayName == "" && fields.About == "" {
			w.WriteHeader(http.StatusNotModified)
			return
		}

		if err := h.profileRepository.Update(r.Context(), fields); err != nil {
			util.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
