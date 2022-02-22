package domain

import (
	"database/sql"
	dtoUser "github.com/pauliusluksys/golang-Reddit/dto/user"
	"gorm.io/gorm"
	"time"
)

type UserGorm struct {
	gorm.Model
	FirstName      string
	LastName       string
	Email          string
	Password       string
	ProfilePicture string
	Country        string
	City           string
}

type User struct {
	ID             uint         `db:"id"`
	CreatedAt      time.Time    `db:"created_at"`
	UpdatedAt      sql.NullTime `db:"updated_at"`
	DeletedAt      sql.NullTime `db:"deleted_at"`
	FirstName      string       `db:"first_name"`
	LastName       string       `db:"last_name"`
	Email          sql.NullString
	Password       string         `db:"passowrd"`
	ProfilePicture sql.NullString `db:"profile_picture"`
	Country        sql.NullString
	City           sql.NullString
}

func FindUser(email string) UserGorm {
	var user UserGorm
	GormDbConnections().Where("email = ?", email).First(&user)
	return user
}

type SignupRequest struct {
	Data dtoUser.SignupRequest
}
type SignupResponse struct {
	Data dtoUser.SignupResponse
}
type LoginRequest struct {
	Data dtoUser.LoginRequest
}
type LoginResponse struct {
	Data dtoUser.LoginResponse
}

func (u UserGorm) LoginResponseToDto() dtoUser.LoginResponse {

	return dtoUser.LoginResponse{
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}
func (s LoginRequest) LoginToDomain() UserGorm {
	sReqData := s.Data

	return UserGorm{
		Email:    sReqData.Email,
		Password: sReqData.Password,
	}
}
func (u UserGorm) SignupResponseToDto() dtoUser.SignupResponse {

	return dtoUser.SignupResponse{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}
}
func (s SignupRequest) ToDomain() UserGorm {
	sReqData := s.Data

	return UserGorm{FirstName: sReqData.FirstName,
		LastName: sReqData.LastName,
		Email:    sReqData.Email,
		Password: sReqData.Password,
	}
}
