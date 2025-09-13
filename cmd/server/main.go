package main

import (
	"log"
	"net/http"
	"os"

	"cashmate-api/config"
	"cashmate-api/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	r.Use(middleware.Logger)

	routes.RegisterRoutes(r)

	log.Printf("Server running on port http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
