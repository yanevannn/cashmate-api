package main

import (
	"log"
	"net/http"
	"os"

	"cashmate-api/config"
	"cashmate-api/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	config.ConnectDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()

	// Middlewares for Logging
	r.Use(middleware.Logger)

	// Middlewares for CORS (Cross-Origin Resource Sharing)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // Use this to allow specific origin hosts
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           30, // Seconds
	}))

	routes.RegisterRoutes(r)

	log.Printf("Server running on port http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
