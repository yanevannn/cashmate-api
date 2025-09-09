package services

import (
    "cashmate-api/models"
    "cashmate-api/repositories"
)

func RegisterUser(user models.User) {
    repositories.CreateUser(user)
}

func ListUsers() []models.User {
    return repositories.GetAllUsers()
}
