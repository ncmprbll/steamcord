package middleware

import (
	"main/backend/internal/models"
	"net/http"
)

func subset(first, second *models.Permissions) bool {
	set := make(map[string]struct{})
	for _, value := range *second {
		set[value] = struct{}{}
	}

	for _, value := range *first {
		if _, found := set[value]; !found {
			return false
		}
	}

	return true
}

func (mw *MiddlewareManager) HasPermissionsMiddleware(required *models.Permissions) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			found, ok := r.Context().Value("user").(*models.User)
			if !ok {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			permissions, err := mw.managementRepository.GetPermissions(r.Context(), found)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			if !subset(required, permissions) {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r.WithContext(r.Context()))
		})
	}
}
