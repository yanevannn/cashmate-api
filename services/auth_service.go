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

	// 6. Send OTP via email using go routine
	// send email in background so that user registration is not delayed
	// we can use worker queue like rabbitmq or beanstalkd for better performance and reliability
	go func(email, OTP, username string) {
		err := utils.SendEmailVerification(email, OTP, username)
		if err != nil {
			fmt.Println("Failed to send email:", err)
		}
	}(user.Email, OTP, user.Username)

	// 7 END
	return nil
}

func ActivateUserService(OTPRequest *models.OTPRequest) error {
	// 1. Fetch user by email
	user, err := repositories.GetUserByEmail(OTPRequest.Email)
	if err != nil {
		return err
	}
	if user == nil {
		return fmt.Errorf("User not found")
	}

	// 2. Check if OTP is valid
	isValid, err := repositories.VerificationOtpIsValid(user.ID, OTPRequest.Code)
	if err != nil {
		return err
	}
	if !isValid {
		return fmt.Errorf("invalid or expired OTP")
	}

	// 3. Mark OTP as used
	if err := repositories.ValidateOTP(user.ID, OTPRequest.Code); err != nil {
		return err
	}

	// 4. Activate user account
	err = repositories.ActivateUser(user.Email)
	if err != nil {
		return err
	}

	return nil
}

func LoginUserService(loginRequest *models.LoginRequest) (*models.LoginTokenResponse, error) {
	// 1. Fetch user by email
	user, err := repositories.GetUserByEmail(strings.ToLower(loginRequest.Email))
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("invalid Credentials")
	}

	// 2. Compare password
	if !utils.CheckPasswordHash(loginRequest.Password, user.Password) {
		return nil, fmt.Errorf("invalid email or passwords")
	}

	// 3. Generate AccessToken JWT & RefreshToken JWT
	accessToken, expiresAt, err := utils.GenerateAccessToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, err
	}
	refreshToken, _, err := utils.GenerateRefreshJWT(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	// 4. Return tokens
	return &models.LoginTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt.Unix(), // Unix timestamp in seconds
	}, nil

}

func RefreshTokenService(refreshToken string) (*models.RefreshTokenResponse, error) {
	// 1. Validate Refresh Token
	claims, err := utils.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %v", err)
	}

	// 2. generate new access token , to get role we need to fetch user from db
	user, err := repositories.GetUserByEmail(claims.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	// 3. Generate new access token
	newAccessToken, expiredAt, err := utils.GenerateAccessToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	// 4. Return new access token
	return &models.RefreshTokenResponse{
		AccessToken: newAccessToken,
		ExpiresAt:   expiredAt.Unix(),
	}, nil

}

func ResendTokenService(userEmail *models.RequestActivateCode) error {
	// 1. Fetch user by email
	user, err := repositories.GetUserByEmail(userEmail.Email)
	if err != nil {
		return err
	}
	if user == nil {
		return fmt.Errorf("User not found")
	}

	// 2. Create OTP
	OTP := utils.GenerateOTP(6)

	// 4. Store OTP in DB
	err = repositories.StoreNewOTP(user.ID, OTP)
	if err != nil {
		return err
	}

	// 5. Send OTP via email using go routine
	go func(email, OTP, username string) {
		err := utils.SendEmailVerification(email, OTP, username)
		if err != nil {
			fmt.Println("Failed to send email:", err)
		}
	}(user.Email, OTP, user.Username)

	return nil
}

func ResendResetPasswordService(userEmail *models.RequestActivateCode) error {
	// 1. Fetch user by email
	user, err := repositories.GetUserByEmail(userEmail.Email)
	if err != nil {
		return err
	}
	if user == nil {
		return fmt.Errorf("User not found")
	}

	// 2. Create OTP
	OTP := utils.GenerateOTP(6)

	// 4. Store OTP in DB
	err = repositories.StoreNewOTP(user.ID, OTP)
	if err != nil {
		return err
	}

	// 5. Send OTP via email using go routine
	go func(email, OTP, username string) {
		err := utils.SendEmailResetPassword(email, OTP, username)
		if err != nil {
			fmt.Println("Failed to send email:", err)
		}
	}(user.Email, OTP, user.Username)

	return nil
}

func ResetPasswordService(resetPasswordRequest *models.ResetPasswordRequest) error {
	// 1. Fetch user by email
	user, err := repositories.GetUserByEmail(resetPasswordRequest.Email)
	if err != nil {
		return err
	}
	if user == nil {
		return fmt.Errorf("User not found")
	}

	// 2. Chcek if OTP is valid
	isValid, err := repositories.VerificationOtpIsValid(user.ID, resetPasswordRequest.Code)
	if err != nil {
		return err
	}
	if !isValid {
		return fmt.Errorf("invalid or expired OTP")
	}

	// 3. Hash new Password 
	hashedpassword, err := utils.HashPassword(resetPasswordRequest.Password)
	if err != nil {
		return err
	}

	// 4. Update password in DB
	err = repositories.UpdateUserPassword(user.Email, user.ID, hashedpassword)
	if err != nil {
		return err
	}

	// 5. Mark OTP as used
	err = repositories.ValidateOTP(user.ID, resetPasswordRequest.Code)
	if err != nil {
		return err
	}

	return nil
}
