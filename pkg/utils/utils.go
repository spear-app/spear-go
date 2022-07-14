package utils

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spear-app/spear-go/pkg/domain/user"
	"golang.org/x/crypto/bcrypt"
)

// ResponseJSON To handle JSON Header and content

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
	claims["user_id"] = user.ID
	claims["name"] = user.Name
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() //Token expires after 1 hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func DeleteJobInternal(db *sql.DB) {
	sqlStatement := `DELETE FROM notifications`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
}
