package main

import (
    "log"
    "net/http"
    "os"

    "github.com/joho/godotenv"
	"cashmate-api/routes"
    "cashmate-api/config"
)

func main() {
    godotenv.Load()

    config.ConnectDB()

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

	routes.RegisterRoutes()
    
    log.Printf("Server running on port http://localhost:%s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
