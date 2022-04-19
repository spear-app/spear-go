package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/spear-app/spear-go/pkg/err"

	"github.com/dgrijalva/jwt-go"
)

// TokenVerifyMiddleware to verify the token before accessing the route
// Middleware
type Claims struct{
	UserId int
}

var (
	ClaimsVar Claims
)
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
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				InterUderID := claims["user_id"]
				strId := fmt.Sprintf("%v", InterUderID)
				userID, err := strconv.Atoi(strId)
				if err!=nil{
					panic(err.Error())
				}
				ClaimsVar.UserId=userID
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

