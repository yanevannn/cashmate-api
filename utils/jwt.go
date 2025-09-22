package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	jwtSecret      = []byte(os.Getenv("JWT_SECRET"))
	accessTokenTTL, _  = time.ParseDuration(os.Getenv("JWT_ACCESS_TTL"))
	refreshTokenTTL, _ = time.ParseDuration(os.Getenv("JWT_REFRESH_TTL"))
)

type JWTClaims struct {
	UserID int `json:"user_id"`
	Email string `json:"email"`
	jwt.RegisteredClaims	
}

func GenerateJWT(userID int, email string) (string, error) {
	expirationTime := time.Now().Add(accessTokenTTL)

	claim := &JWTClaims{
		UserID: userID,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (*JWTClaims, error) {
	claim := &JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claim, nil
}

func GetAccessTokenTTL() time.Duration {
	return accessTokenTTL
}

func GetRefreshTokenTTL() time.Duration {
	return refreshTokenTTL
}