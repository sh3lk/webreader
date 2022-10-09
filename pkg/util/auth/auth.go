package auth

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
	"webreader/ent/schema/ulid"
)

var secret = []byte("secret")

type jwtClaims struct {
	Id ulid.ID `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(userId ulid.ID) (string, error) {
	// Set custom claims
	claims := &jwtClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	// Generate encoded token and send it as response.
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
}

func ParseToken(s string) (ulid.ID, error) {
	token, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if token == nil || err != nil {
		return "", err
	}
	claims := token.Claims.(jwt.MapClaims)

	return ulid.ID(claims["id"].(string)), err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
