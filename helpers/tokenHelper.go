package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SECRET_KEY []byte

type SignDetail struct {
	Email 	string
	User_id	string
	jwt.RegisteredClaims
}

func GenerateToken(email, userID string) (string, string, error) {
	claims := &SignDetail{
		Email: email,
		User_id: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(SECRET_KEY)
	if err != nil {
		return "", "", nil
	}

	refreshClaims := &SignDetail{
		Email: email,
		User_id: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		},
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(SECRET_KEY)
	if err != nil {
		return "", "", nil
	}
	

	return token, refreshToken, nil
}

func ValidateToken(signedToken string) (*SignDetail, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignDetail{},
		func(t *jwt.Token) (interface{}, error) {
			return SECRET_KEY, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*SignDetail)
	if !ok {
		return nil, err
	}

	return claims, nil
}