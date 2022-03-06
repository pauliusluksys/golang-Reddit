package services

import (
	"github.com/hashicorp/go-hclog"
	"github.com/pauliusluksys/golang-Reddit/domain"
	dtoUser "github.com/pauliusluksys/golang-Reddit/dto/user"
	"github.com/pauliusluksys/golang-Reddit/errs"
	"github.com/pauliusluksys/golang-Reddit/utils"
	"github.com/pauliusluksys/golang-Reddit/validation"
	"gorm.io/gorm"
	"time"
)

type UserService struct {
	UserServiceGorm domain.UserGormRepo
	Logger          hclog.Logger
}

func NewUserService(usqlxRepo domain.UserGormRepo, logger hclog.Logger) UserService {
	return UserService{UserServiceGorm: usqlxRepo, Logger: logger}
}
func (s UserService) Login(request dtoUser.LoginRequest) (*dtoUser.LoginResponse, *errs.AppError) {

	err := validation.Login(request)
	if err != nil {
		return nil, err
	}

	userGorm, rowsAffected, err := s.UserServiceGorm.FindUser(request.Email)
	if err != nil {
		return nil, err
	}
	if int(*rowsAffected) < 1 {
		return nil, errs.NewNotFoundError("Invalid Email, Please Signup!")

	}

	isEqual := utils.DoPasswordsMatch(userGorm.Password, request.Password)
	if !isEqual {
		return nil, errs.NewNotFoundError("Invalid Credentials!")

	}
	payload := utils.Payload{
		Email: userGorm.Email,
	}
	token, err := utils.GenerateJwtToken(payload)
	if err != nil {
		return nil, err
	}
	loginResponse := userGorm.LoginResponseToDto(*token)

	return &loginResponse, nil
}
func (s UserService) Signup(request dtoUser.SignupRequest) (*dtoUser.SignupResponse, *errs.AppError) {
	err := validation.Signup(request)
	if err != nil {
		return nil, err
	}
	newUser := domain.UserGorm{
		Model:     gorm.Model{CreatedAt: time.Now()},
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
	}
	user, _, err := s.UserServiceGorm.CreateUser(newUser)
	if err != nil {
		return nil, err
	}
	payload := utils.Payload{
		Email: user.Email,
	}
	token, err := utils.GenerateJwtToken(payload)
	if err != nil {
		return nil, err
	}
	signupResponse := user.SignupResponseToDto(*token)

	return &signupResponse, nil
}
