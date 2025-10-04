package middlewares

import (
	"cashmate-api/utils"
	"context"
	"net/http"
	"strings"
)

type contextKey string

const contextKeyClaims contextKey = "claims"


func AuthMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header from the request
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.ResError(w, http.StatusUnauthorized, "Missing Authorization header")
			return
		}

		// Clean the token by removing "Bearer " prefix if present
		token := strings.TrimPrefix(authHeader, "Bearer ")
		token = strings.TrimSpace(token)
		if token == "" {
			utils.ResError(w, http.StatusUnauthorized, "Missing token")
			return
		}

		// Validate the Token
		claims, err := utils.ValidateAccessToken(token)
		if err != nil {
			utils.ResError(w, http.StatusUnauthorized, "Invalid token: "+err.Error())
			return
		}

		// Add user id and role to the request context
		ctx := context.WithValue(r.Context(), contextKeyClaims, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetClaimsFromCtx(r *http.Request) (*utils.AccessTokenClaims, bool) {
	claims, ok := r.Context().Value(contextKeyClaims).(*utils.AccessTokenClaims)
	return claims, ok
}