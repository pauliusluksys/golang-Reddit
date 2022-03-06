package domain

import (
	"database/sql"
	"github.com/hashicorp/go-hclog"
	dtoUser "github.com/pauliusluksys/golang-Reddit/dto/user"
	"github.com/pauliusluksys/golang-Reddit/errs"
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
type UserGormRepo struct {
	Client *gorm.DB
	Logger hclog.Logger
}

func NewUserGormRepoDb(gormDb *gorm.DB, logger hclog.Logger) UserGormRepo {
	return UserGormRepo{Client: gormDb, Logger: logger}
}
func (ugdb UserGormRepo) FindUser(email string) (*UserGorm, *int64, *errs.AppError) {
	var user UserGorm
	result := ugdb.Client.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, nil, errs.NewUnexpectError("Gorm has failed while fetching user" + result.Error.Error())
	}
	return &user, &result.RowsAffected, nil
}
func (ugdb UserGormRepo) CreateUser(user UserGorm) (*UserGorm, *int64, *errs.AppError) {
	result := ugdb.Client.Create(&user)
	if result.Error != nil {
		return nil, nil, errs.NewUnexpectError("Gorm has failed while fetching user" + result.Error.Error())
	}
	return &user, &result.RowsAffected, nil
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

func (u UserGorm) LoginResponseToDto(token string) dtoUser.LoginResponse {

	return dtoUser.LoginResponse{
		User: dtoUser.UserResponse{
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
		},
		Token: dtoUser.TokenResponse{
			Token: token,
		},
	}

}
func (s LoginRequest) LoginToDomain() UserGorm {
	sReqData := s.Data

	return UserGorm{
		Email:    sReqData.Email,
		Password: sReqData.Password,
	}
}
func (u UserGorm) SignupResponseToDto(token string) dtoUser.SignupResponse {

	return dtoUser.SignupResponse{
		UserResponse: dtoUser.UserResponse{
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
		},
		TokenResponse: dtoUser.TokenResponse{
			Token: token,
		},
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
