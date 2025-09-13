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

func GetAllUsersService() ([]models.PublicUser, error) {
	users , err := repositories.GetAllUsers()
	if err != nil {
		return nil, err
	}
	// Extract data to exclude passwords
	var publicUsers []models.PublicUser
	for _, user := range users {
		publicUsers = append(publicUsers, models.PublicUser{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		})
	}
	return publicUsers, nil
}

func DeleteUserService(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid user ID")
	}
	err := repositories.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}