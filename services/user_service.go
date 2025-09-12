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

func GetUserByIDService(id int) (*models.User, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid user ID")
	}

	user, err := repositories.GetuserByID(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}