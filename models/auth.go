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
	IsActive bool `json:"is_active"`
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

type OTPRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func (r OTPRequest) Validate () error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Code, validation.Required, validation.Length(6, 6)),
	)
}

type RequestActivateCode struct {
	Email string `json:"email"`
}

func (r RequestActivateCode) Validate () error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.Email, validation.Required, is.Email),
	)
}

type ResetPasswordRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
	Password string `json:"password"`
}

func (r ResetPasswordRequest) Validate () error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Code, validation.Required, validation.Length(6, 6)),
		validation.Field(&r.Password, validation.Required, validation.Length(6, 100)),
	)
}

