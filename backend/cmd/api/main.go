package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/redis/go-redis/v9"

	authDelivery "main/backend/internal/auth/delivery/http"
	productsDelivery "main/backend/internal/products/delivery/http"
	cartDelivery "main/backend/internal/cart/delivery/http"
	languageDelivery "main/backend/internal/language/delivery/http"
	profileDelivery "main/backend/internal/profile/delivery/http"
	authRepository "main/backend/internal/auth/postgres"
	productsRepository "main/backend/internal/products/postgres"
	sessionRepository "main/backend/internal/session/redis"
	cartRepository "main/backend/internal/cart/postgres"
	languageRepository "main/backend/internal/language/postgres"
	profileRepository "main/backend/internal/profile/postgres"

	mw "main/backend/internal/middleware"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func main() {
	url := "postgres://postgres:password@localhost/postgres"
	port := "3000"

	database, err := sqlx.Open("pgx", url)

	if err != nil {
		panic(err)
	}

	err = database.Ping()

	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err = rdb.Set(context.TODO(), "key", "value", time.Second * 5).Err()
	if err != nil {
		panic(err)
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
	sessionRedis := sessionRepository.New(rdb)
	productsPostgres := productsRepository.New(database)
	cartPostgres := cartRepository.New(database)
	languagePostgres := languageRepository.New(database)
	profilePostgres := profileRepository.New(database)

	// Middleware Manager
	manager := mw.NewMiddlewareManager(authPostgres, sessionRedis)

	// Handlers
	authHandlers := authDelivery.NewAuthHandlers(authPostgres, sessionRedis)
	productsHandlers := productsDelivery.NewAuthHandlers(productsPostgres)
	cartHandlers := cartDelivery.NewAuthHandlers(cartPostgres, authPostgres, sessionRedis)
	languageHandlers := languageDelivery.NewAuthHandlers(languagePostgres)
	profileHandlers := profileDelivery.NewAuthHandlers(profilePostgres)

	// Routers
	authRouter := authDelivery.NewRouter(authHandlers, manager)
	productsRouter := productsDelivery.NewRouter(productsHandlers, manager)
	cartRouter := cartDelivery.NewRouter(cartHandlers, manager)
	languageRouter := languageDelivery.NewRouter(languageHandlers)
	profileRouter := profileDelivery.NewRouter(profileHandlers, manager)

	r.Mount("/auth", authRouter)
	r.Mount("/products", productsRouter)
	r.Mount("/cart", cartRouter)
	r.Mount("/locales", languageRouter)
	r.Mount("/profile", profileRouter)

	http.ListenAndServe(":"+port, r)
}
