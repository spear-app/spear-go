package utils

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator"

	"github.com/spear-app/spear-go/pkg/domain/user"
	"github.com/spear-app/spear-go/pkg/err"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// ResponseJSON To handle JSON Header and content

// ValidateInputs to validate user inputs
func ValidateInputs(w http.ResponseWriter, err error) {
	for _, e := range err.(validator.ValidationErrors) {
		switch e.ActualTag() {
		case "gte":
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errs.NewResponse(e.Field()+" is less than 8 characters", http.StatusBadRequest))

		case "email":
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errs.NewResponse(e.Field()+" is not a valid email", http.StatusBadRequest))
		case "len":
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errs.NewResponse(e.Field()+" is not a valid phone number", http.StatusBadRequest))
		case "alpha":
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errs.NewResponse(e.Field()+" is not a valid name", http.StatusBadRequest))

		default:
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errs.NewResponse(e.Field()+" is "+e.ActualTag(), http.StatusBadRequest))
		}
		return
	}
}

// HashPassword To hash passwords
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// CheckPasswordHash To compare hash password with the one in DB
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateToken to generate token when user is signed in
func GenerateToken(user user.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["sub"] = user.ID
	claims["name"] = user.Name
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}
