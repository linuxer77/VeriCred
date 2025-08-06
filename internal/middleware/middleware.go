package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"vericred/pkg"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		fmt.Println("Inside the Auth Middleware")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}
	
		log.Println("AuthHeader: ", authHeader)

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			http.Error(w, "Invalid auth format", http.StatusUnauthorized)
			return
		}

		log.Println("BearerToken: ", bearerToken)

		token := bearerToken[1]
		address, err := pkg.VerifyToken(token)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		ctx := context.WithValue(r.Context(), "metamaskAddress", address)        
        next.ServeHTTP(w, r.WithContext(ctx))

	})
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Missing auth header.")
		return
	}
	tokenString = tokenString[len("Bearer "):]

    _, err := pkg.VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}

	fmt.Fprint(w, "welcome fucker.")
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s %s\n", r.Method, r.RemoteAddr, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

