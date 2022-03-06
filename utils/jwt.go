package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/pauliusluksys/golang-Reddit/errs"
	"time"
)

type Payload struct {
	Email string
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var JWT_SECRET string

func GenerateJwtToken(payload Payload) (*string, *errs.AppError) {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		return nil, errs.NewUnexpectError("failed to read env file: " + err.Error())
	}
	if JWT_SECRET = myEnv["JWT_SECRET"]; JWT_SECRET == "" {
		return nil, errs.NewUnexpectError("[ ERROR ] JWT_SECRET environment variable not provided!\n")
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
		return nil, errs.NewUnexpectError("error while trying to sign new token: " + err.Error())

	}

	return &SignedToken, nil
}

func VerifyJwtToken(strToken string) (*Claims, *errs.AppError) {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		return nil, errs.NewUnexpectError("failed to read env file: " + err.Error())
	}

	if JWT_SECRET = myEnv["JWT_SECRET"]; JWT_SECRET == "" {
		return nil, errs.NewUnexpectError("[ ERROR ] JWT_SECRET environment variable not provided!\n")
	}

	key := []byte(JWT_SECRET)

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(strToken, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {

			return claims, errs.NewForbiddenError("invalid token signature")
		} else {
			return nil, errs.NewUnexpectError(err.Error())
		}
	}

	if !token.Valid {
		return claims, errs.NewValidationError("invalid token")
	}

	return claims, nil
}
