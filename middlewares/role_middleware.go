package middlewares

import (
	"cashmate-api/utils"
	"net/http"
)

func RoleMiddleware(allowedRoles ...string) func(http.Handler) http.Handler {

	// 1. Convert allowedRoles slice to a map for efficient lookup
	roleMap := map[string]struct{}{}
	for _, role := range allowedRoles {
		roleMap[role] = struct{}{}
	}

	// 2. Return the middleware function
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 3. Extract claims from context
			claims, ok := GetClaimsFromCtx(r)
			if !ok {
				utils.ResError(w, http.StatusUnauthorized, "Unauthorized: missing or invalid token claims")
				return
			}

			// 4. Get the user's role from claims
			userRole := claims.Role

			// 5. Check if the user's role is in the allowed roles
			if _, roleAllowed := roleMap[userRole]; !roleAllowed {
				utils.ResError(w, http.StatusForbidden, "You do not have permission to access this resource")
				return
			}

			// 6. Proceed to the next handler if role is allowed
			next.ServeHTTP(w, r)
		})
	}
}
