package validation

import (
	dtoUser "github.com/pauliusluksys/golang-Reddit/dto/user"
	"github.com/pauliusluksys/golang-Reddit/errs"
)

func Login(request dtoUser.LoginRequest) *errs.AppError {
	if len(request.Email) < 3 {
		return errs.NewValidationError("Email Cannot be shorter than 3 symbols")
	}

	if len(request.Password) < 3 {
		return errs.NewValidationError("Password cannot be shorter than 3 symbols")
	}
	return nil
}
func Signup(request dtoUser.SignupRequest) *errs.AppError {
	if len(request.Email) < 3 {
		return errs.NewValidationError("Email cannot be shorter than 3 symbols")
	}

	if len(request.Password) < 3 {
		return errs.NewValidationError("Password cannot be shorter than 3 symbols")
	}
	if len(request.FirstName) < 3 {
		return errs.NewValidationError("First name cannot be shorter than 3 symbols")
	}
	if len(request.LastName) < 3 {
		return errs.NewValidationError("Last name cannot be shorter than 3 symbols")
	}
	return nil
}
