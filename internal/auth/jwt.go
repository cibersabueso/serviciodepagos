package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JwtKey es la clave secreta usada para firmar los tokens JWT.
var JwtKey = []byte("})VV}40Nl#0n")

type Claims struct {
	Username string `json:"pruebas"`
	jwt.StandardClaims
}

func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
