package middleware

import (
	"main/backend/internal/auth"
	"main/backend/internal/session"
)

type MiddlewareManager struct {
	authRepository auth.Repository
	sessionRepository session.Repository
}

func NewMiddlewareManager(aR auth.Repository, sR session.Repository) *MiddlewareManager {
	return &MiddlewareManager{aR, sR}
}