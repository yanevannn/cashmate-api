package services

import (
	"cashmate-api/models"
	"cashmate-api/repositories"
	"fmt"
)

func CreateUserService(user *models.User) error {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return fmt.Errorf("all fields are required")

	}
	return repositories.CreateUser(user)
}
