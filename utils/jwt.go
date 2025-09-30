package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ⚠️ Do NOT read ENV directly in global vars,
// because it runs when the package is imported,
// before godotenv.Load() is called in main().
// That makes os.Getenv("JWT_SECRET") return empty.

// ✅ Use functions instead, so ENV values
// are read only when needed (after godotenv.Load()).
func getJWTSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func getAccessTokenTTL() time.Duration {
	d, _ := time.ParseDuration(os.Getenv("JWT_ACCESS_TTL"))
	return d
}

func getRefreshTokenTTL() time.Duration {
	d, _ := time.ParseDuration(os.Getenv("JWT_REFRESH_TTL"))
	return d
}

type AccessTokenClaims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int, email string, role string) (string, time.Time, error) {
	expirationTime := time.Now().Add(getAccessTokenTTL())
	jwtSecret := getJWTSecret()

	claim := &AccessTokenClaims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "cashmate-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // token : payload raw
	signin, err := token.SignedString(jwtSecret)              // signed token : payload encoded
	return signin, expirationTime, err
}

func GenerateRefreshJWT(userID int, email string) (string, time.Time, error) {
	jwtSecret := getJWTSecret()
	expirationTime := time.Now().Add(getRefreshTokenTTL())

	claim := &RefreshTokenClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "cashmate-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // token : payload raw
	signin, err := token.SignedString(jwtSecret)              // signed token : payload encoded
	return signin, expirationTime, err
}

func ValidateAccessToken(tokenString string) (*AccessTokenClaims, error) {
	jwtSecret := getJWTSecret()
	claim := &AccessTokenClaims{}
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

func ValidateRefreshToken(tokenString string) (*RefreshTokenClaims, error) {
	jwtSecret := getJWTSecret()
	claim := &RefreshTokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err // Check for parsing errors (e.g., malformed token, signature invalid, token expired
	}

	if !token.Valid {
		return nil, errors.New("invalid token") // Check if token is valid
	}

	return claim, nil
}
