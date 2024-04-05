package main

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	authDelivery "main/backend/internal/auth/delivery/http"
	authRepository "main/backend/internal/auth/postgres"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	url := "postgres://postgres:password@localhost/postgres"
	port := "3000"

	database, err := sql.Open("pgx", url)

	if err != nil {
		return
	}

	err = database.Ping()

	if err != nil {
		return
	}

	r := chi.NewRouter()
	r.Use(middleware.DefaultLogger)

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Repositories
	authPostgres := authRepository.New(database)

	// Handlers
	authHandlers := authDelivery.NewAuthHandlers(authPostgres)

	// Routers
	authRouter := authDelivery.NewRouter(authHandlers)

	r.Mount("/auth", authRouter)

	http.ListenAndServe(":" + port, r)
}
