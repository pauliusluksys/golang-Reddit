package dtoUser

import "time"

type SignupRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
type SignupResponse struct {
	UserResponse  `json:"user_response"`
	TokenResponse `json:"token_response"`
}
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginResponse struct {
	User  UserResponse  `json:"user_response"`
	Token TokenResponse `json:"token_response"`
}
type VerificationData struct {
	Email     string    `json:"email" validate:"required" sql:"email"`
	Code      string    `json:"code" validate:"required" sql:"code"`
	ExpiresAt time.Time `json:"expiresat" sql:"expiresat"`
	//Type      VerificationDataType `json:"type" sql:"type"`
}
type UserResponse struct {
	FirstName     string     `json:"first_name"`
	LastName      string     `json:"last_name"`
	Email         string     `json:"email"`
	EmailVerified *time.Time `json:"email_verified"`
}

type TokenResponse struct {
	Token string
}
