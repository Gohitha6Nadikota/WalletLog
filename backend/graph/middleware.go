
package graph

import (
	"context"
	"net/http"
	"strings"
	"log"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next *handler.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		log.Println("Authorization Header:", authHeader)

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			writeGraphQLError(w, http.StatusUnauthorized, "Missing or invalid Authorization header")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		secret := []byte("your-secret-key")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if err != nil || !token.Valid {
			writeGraphQLError(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			writeGraphQLError(w, http.StatusUnauthorized, "Invalid token claims")
			return
		}

		userID, ok := claims["userID"].(string)
		if !ok {
			writeGraphQLError(w, http.StatusUnauthorized, "Missing user ID in token")
			return
		}

	
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func writeGraphQLError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(`{"errors":[{"message":"` + message + `"}]}`))
}
