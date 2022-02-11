package middleware

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/spear-app/spear-go/pkg/err"

	"github.com/dgrijalva/jwt-go"
)

// TokenVerifyMiddleware to verify the token before accessing the route
// Middleware
func TokenVerifyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) == 2 {
			authToken := bearerToken[1]
			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, json.NewEncoder(w).Encode("There was an error!")
				}
				return []byte(os.Getenv("API_SECRET")), nil
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(errs.NewResponse(err.Error(), http.StatusUnauthorized))
				return
			}
			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(errs.NewResponse(err.Error(), http.StatusUnauthorized))
				return
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errs.NewResponse(errs.ErrInvalidToken.Error(), http.StatusUnauthorized))
			return
		}
	})
}
