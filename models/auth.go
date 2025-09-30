package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)


type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Role string `json:"role"`
}

type RegisterUser struct {
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6	"`
}

func (r RegisterUser) Validate () error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.Username, validation.Required),
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(6, 100)),
	)
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func (r LoginRequest) Validate () error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(6, 100)),
	)
}

type LoginTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (r RefreshTokenRequest) Validate () error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.RefreshToken, validation.Required),
	)
}

type RefreshTokenResponse struct {
    AccessToken string `json:"access_token"`
    ExpiresAt   int64  `json:"expires_at"`
}

