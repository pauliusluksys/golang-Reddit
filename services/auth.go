package services

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/pauliusluksys/golang-Reddit/domain"
	"github.com/pauliusluksys/golang-Reddit/errs"
	"github.com/pauliusluksys/golang-Reddit/utils"
	"strings"
)

type AuthService struct {
	repo   domain.UserGormRepo
	logger hclog.Logger
}

// NewAuthService returns a new instance of the auth service
func NewAuthService(userGormRepo domain.UserGormRepo, logger hclog.Logger) AuthService {
	return AuthService{userGormRepo, logger}
}

func (s AuthService) CheckAuth(authHeader string) (*domain.UserGorm, *errs.AppError) {
	bearerToken := strings.Split(authHeader, " ")
	fmt.Println("bearer token:    ", bearerToken)
	fmt.Println()
	if len(bearerToken) < 2 {
		return nil, errs.NewForbiddenError("Token not provided")
	}

	token := bearerToken[1]
	s.logger.Debug(token)
	claims, err := utils.VerifyJwtToken(token)
	if err != nil {
		return nil, err
	}
	user, rows, err := s.repo.FindUser(claims.Email)
	s.logger.Debug("retrieved user data from repo: ", user)
	if err != nil {
		return nil, err
	} else if int(*rows) < 1 {
		return nil, errs.NewForbiddenError("User does not exist")
	}
	return user, nil

}
