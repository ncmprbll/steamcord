package middleware

import (
	"main/backend/internal/auth"
	"main/backend/internal/management"
	"main/backend/internal/session"
)

type MiddlewareManager struct {
	authRepository auth.Repository
	sessionRepository session.Repository
	managementRepository management.Repository
}

func NewMiddlewareManager(aR auth.Repository, sR session.Repository, mR management.Repository) *MiddlewareManager {
	return &MiddlewareManager{aR, sR, mR}
}