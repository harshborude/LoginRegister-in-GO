package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var accessSecret = []byte("access-secret")
var refreshSecret = []byte("refresh-secret")

type Claims struct {
	UserID uint
	Role   string
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID uint, role string) (string, error) {

	expiration := time.Now().Add(15 * time.Minute)

	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(accessSecret)
}

func GenerateRefreshToken(userID uint) (string, error) {

	expiration := time.Now().Add(7 * 24 * time.Hour)

	claims := jwt.RegisteredClaims{
		Subject: fmt.Sprint(userID),
		ExpiresAt: jwt.NewNumericDate(expiration),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(refreshSecret)
}

func ValidateAccessToken(tokenString string) (*Claims, error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return accessSecret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func ValidateRefreshToken(tokenString string) (*jwt.RegisteredClaims, error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return refreshSecret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)

	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}