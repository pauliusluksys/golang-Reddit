package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Payload struct {
	Email string
}

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Id       uint   `json:"id"`
	jwt.RegisteredClaims
}

var JWT_SECRET string

func GenerateJwtToken(payload Payload) (string, error) {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if JWT_SECRET = myEnv["JWT_SECRET"]; JWT_SECRET == "" {
		log.Fatal("[ ERROR ] JWT_SECRET environment variable not provided!\n")
	}

	key := []byte(JWT_SECRET)

	expirationTime := time.Now().Add(7 * 24 * 60 * time.Minute)

	claims := &Claims{
		Email: payload.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expirationTime},
		},
	}

	UnsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	SignedToken, err := UnsignedToken.SignedString(key)
	if err != nil {
		return "", err
	}

	return SignedToken, nil
}

func VerifyJwtToken(strToken string) (*Claims, error) {
	if JWT_SECRET = os.Getenv("JWT_SECRET"); JWT_SECRET == "" {
		log.Fatal("[ ERROR ] JWT_SECRET environment variable not provided!\n")
	}

	key := []byte(JWT_SECRET)

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(strToken, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return claims, fmt.Errorf("invalid token signature")
		}
	}

	if !token.Valid {
		return claims, fmt.Errorf("invalid token")
	}

	return claims, nil
}
