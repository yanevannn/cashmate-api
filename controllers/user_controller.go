package controllers

import (
    "cashmate-api/models"
    "cashmate-api/services"
    "encoding/json"
    "net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    switch r.Method {
    case http.MethodGet:
        users := services.ListUsers()
        json.NewEncoder(w).Encode(users)

    case http.MethodPost:
        var user models.User
        json.NewDecoder(r.Body).Decode(&user)
        services.RegisterUser(user)
        json.NewEncoder(w).Encode(map[string]string{"message": "User created"})

    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
        json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
    }
}
