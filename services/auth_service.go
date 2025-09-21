package services

import (
	"cashmate-api/models"
	"cashmate-api/repositories"
	"cashmate-api/utils"
	"fmt"
	"strings"
)

func RegisterUserService(user *models.RegisterUser) error {
	// 1. Check existing user
	existingUser, err := repositories.GetUserByEmail(user.Email)
	if err != nil {
		return err
	} else if existingUser != nil {
		return fmt.Errorf("user with this email already exists")
	}

	// 2. Normalize email to lowercase
	user.Email = strings.ToLower(user.Email)

	// 3. Hash Password
	hashedpassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedpassword

	// 5. Create OTP
	OTP := utils.GenerateOTP(6)

	// Create new user
	err = repositories.CreateUser(user, OTP)
	if err != nil {
		return err
	}

	// 6. Send OTP via email
	// COMING SOON

	return nil
}
