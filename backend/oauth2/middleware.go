package oauth2

import (
    "net/http"
    "strings"
)

func OAuth2Middleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Check if a valid OAuth2 token is present in the request headers
        bearerToken := r.Header.Get("Authorization")
        if bearerToken != "" {
            tokenParts := strings.Split(bearerToken, " ")
            if len(tokenParts) == 2 && tokenParts[0] == "Bearer" {
                // Validate the token here (e.g., check expiration, scopes, etc.)
                // If valid, call the next handler
                next(w, r)
                return
            }
        }

        http.Error(w, "Unauthorized", http.StatusUnauthorized)
    }
}
