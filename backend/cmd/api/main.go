package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/webhook"

	authDelivery "main/backend/internal/auth/delivery/http"
	authRepository "main/backend/internal/auth/postgres"
	cartDelivery "main/backend/internal/cart/delivery/http"
	cartRepository "main/backend/internal/cart/postgres"
	languageDelivery "main/backend/internal/language/delivery/http"
	languageRepository "main/backend/internal/language/postgres"
	managementDelivery "main/backend/internal/management/delivery/http"
	managementRepository "main/backend/internal/management/postgres"
	productsDelivery "main/backend/internal/products/delivery/http"
	productsRepository "main/backend/internal/products/postgres"
	profileDelivery "main/backend/internal/profile/delivery/http"
	profileRepository "main/backend/internal/profile/postgres"
	sessionRepository "main/backend/internal/session/redis"

	mw "main/backend/internal/middleware"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	webhookSecret := os.Getenv("WEBHOOK_SECRET")
	postgresSource := os.Getenv("SERVER_POSTGRES_SOURCE")
	redisSource := os.Getenv("SERVER_REDIS_SOURCE")
	port := "3000"

	database, err := sqlx.Open("pgx", postgresSource)

	if err != nil {
		panic(err)
	}

	err = database.Ping()

	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisSource,
		Password: "",
		DB:       0,
	})

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
	managementPostgres := managementRepository.New(database)

	// Middleware Manager
	manager := mw.NewMiddlewareManager(authPostgres, sessionRedis, managementPostgres)

	// Handlers
	authHandlers := authDelivery.NewAuthHandlers(authPostgres, sessionRedis, profilePostgres)
	productsHandlers := productsDelivery.NewAuthHandlers(productsPostgres)
	cartHandlers := cartDelivery.NewAuthHandlers(cartPostgres, authPostgres, sessionRedis)
	languageHandlers := languageDelivery.NewAuthHandlers(languagePostgres)
	profileHandlers := profileDelivery.NewAuthHandlers(sessionRedis, profilePostgres)
	managementHandlers := managementDelivery.NewManagementHandlers(managementPostgres, sessionRedis)

	// Routers
	authRouter := authDelivery.NewRouter(authHandlers, manager)
	productsRouter := productsDelivery.NewRouter(productsHandlers, manager)
	cartRouter := cartDelivery.NewRouter(cartHandlers, manager)
	languageRouter := languageDelivery.NewRouter(languageHandlers)
	profileRouter := profileDelivery.NewRouter(profileHandlers, manager)
	managementRouter := managementDelivery.NewRouter(managementHandlers, manager)

	r.Route("/api", func(r chi.Router) {
		r.Mount("/auth", authRouter)
		r.Mount("/products", productsRouter)
		r.Mount("/cart", cartRouter)
		r.Mount("/locales", languageRouter)
		r.Mount("/profile", profileRouter)
		r.Mount("/management", managementRouter)
	})

	r.Post("/webhook", func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 2<<15)

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		event, err := webhook.ConstructEvent(body, r.Header.Get("Stripe-Signature"), webhookSecret)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if event.Type == "checkout.session.completed" {
			var session stripe.CheckoutSession
			err := json.Unmarshal(event.Data.Raw, &session)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if _, err := database.ExecContext(r.Context(), "UPDATE users SET balance = balance + $1 WHERE id = $2", session.AmountTotal/100, session.Metadata["user_id"]); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
		}

		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":"+port, r)
}
