package main

import (
    "log"
    "net/http"
    "os"

    "github.com/joho/godotenv"
	"cashmate-api/routes"
)

func main() {
    godotenv.Load()

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

	routes.RegisterRoutes()
    
    log.Println("Server running on port", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
