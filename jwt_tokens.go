package authboss

import(
	"github.com/dgrijalva/jwt-go"
	"time"
)

var mySigningKey = []byte("secret")

func NewJWTToken(email string) string{
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"name": email, "exp": time.Now().Add(time.Hour * 24).Unix()})
	tokenString, _ := token.SignedString(mySigningKey)

	return tokenString
}
