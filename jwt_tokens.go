package authboss

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func NewJWTToken(email string, secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"name": email, "exp": time.Now().Add(time.Hour * 24).Unix()})
	tokenString, _ := token.SignedString([]byte(secret))

	return tokenString
}
