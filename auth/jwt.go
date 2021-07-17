package auth

import (
	"github.com/dgrijalva/jwt-go"
	"go-picbed/model"
	"time"
)

var JwtKey = []byte("114514")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func TokenRelease(user model.User) (string, error) {
	ExpTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ExpTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "hibana",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	TokenString, err := token.SignedString(JwtKey)

	if err != nil {
		return "", err
	}

	return TokenString, nil
}

func TokenParse(TokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(TokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	return token, claims, err

}
