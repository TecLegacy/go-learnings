package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/teclegacy/rss-aggregator/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port String not found")
	}

	// Postgres DB string
	dbString := os.Getenv("DB_URL")
	if dbString == "" {
		log.Fatal("Postgres String not found")
	}

	// Connection to postgres server
	conn, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal("Failed to connect to DB", err)
	}
	query := database.New(conn)
	apiCfg := apiConfig{
		DB: query,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(
		cors.Options{
			AllowedOrigins: []string{
				"https://*",
				"http://*",
			},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowCredentials: false,
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			// AllowedHeaders: []string{"*"},
			ExposedHeaders: []string{"Link"},
			MaxAge:         300,
		},
	))

	v1Router := chi.NewRouter()

	v1Router.HandleFunc("/healthz", handlerHealth)

	v1Router.Get("/error", handlerError)

	// Create user -> response with authenticated API_KEY
	v1Router.Post("/user", apiCfg.createUserHandler)
	// Response with current user if authenticated with API else Not authorized
	v1Router.Get("/user", apiCfg.authMiddleware(apiCfg.getUserHandler))

	// Authenticated user creating a feed
	v1Router.Post("/feed", apiCfg.authMiddleware(apiCfg.handlerCreateFeed))
	// To Get all feeds in our db
	v1Router.Get("/feed", apiCfg.handlerGetAllFeeds)

	// User likes a particular feed
	v1Router.Post("/feed_follows", apiCfg.authMiddleware(apiCfg.handleFeedFollowsRequest))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}

	fmt.Printf("Server Started and listening %v", portString)

	// Server Start
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
