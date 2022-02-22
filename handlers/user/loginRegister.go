package userHandler

import (
	"encoding/json"
	"github.com/pauliusluksys/golang-Reddit/domain"
	dtoUser "github.com/pauliusluksys/golang-Reddit/dto/user"
	"github.com/pauliusluksys/golang-Reddit/utils"
	"gorm.io/gorm"
	"net/http"
)

var errors = utils.CustomError{}

type SignupResponseOutput struct {
	User  dtoUser.SignupResponse
	Token string
}
type LoginResponseOutput struct {
	User  dtoUser.LoginResponse
	Token string
}

func UserLogin(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loginRequest := dtoUser.LoginRequest{}

		err := json.NewDecoder(r.Body).Decode(&loginRequest)
		if err != nil {
			return
		}
		if len(loginRequest.Email) < 3 {
			errors.ApiError(w, http.StatusBadRequest, "Invalid Email!")
			return
		}

		if len(loginRequest.Password) < 3 {
			errors.ApiError(w, http.StatusBadRequest, "Invalid Password!")
			return
		}
		//dtoLoginReq := domain.LoginRequest{Data: loginRequest}
		//domainUser := dtoLoginReq.LoginToDomain()
		var userGorm domain.UserGorm

		if results := db.Where("email = ?", loginRequest.Email).First(&userGorm); results.Error != nil || results.RowsAffected < 1 {
			errors.ApiError(w, http.StatusNotFound, "Invalid Email, Please Signup!")
			return
		}
		isEqual := utils.DoPasswordsMatch(userGorm.Password, loginRequest.Password)
		if !isEqual {
			errors.ApiError(w, http.StatusNotFound, "Invalid Credentials!")
			return
		}
		loginResponse := userGorm.LoginResponseToDto()
		payload := utils.Payload{
			Email: userGorm.Email,
		}

		token, err := utils.GenerateJwtToken(payload)
		if err != nil {
			errors.ApiError(w, http.StatusInternalServerError, "Failed To Generate New JWT Token!")
			return
		}

		utils.RespondWithJSON(w, LoginResponseOutput{
			Token: token,
			User:  loginResponse,
		})

	}
}
func UserSignup(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		signupRequest := dtoUser.SignupRequest{}
		json.NewDecoder(r.Body).Decode(&signupRequest)

		if len(signupRequest.FirstName) < 3 {
			errors.ApiError(w, http.StatusBadRequest, "Username should be at least 3 characters long!")
			return
		}

		if len(signupRequest.LastName) < 3 {
			errors.ApiError(w, http.StatusBadRequest, "Name should be at least 3 characters long!")
			return
		}

		if len(signupRequest.Email) < 3 {
			errors.ApiError(w, http.StatusBadRequest, "Email should be at least 3 characters long!")
			return
		}

		if len(signupRequest.Password) < 3 {
			errors.ApiError(w, http.StatusBadRequest, "Password should be at least 3 characters long!")
			return
		}
		dtoSignupReq := domain.SignupRequest{Data: signupRequest}
		domainUser := dtoSignupReq.ToDomain()
		signupResponse := domainUser.SignupResponseToDto()
		domainUser.Password = utils.EncryptPassword(domainUser.Password)

		if result := db.Create(&domainUser); result.Error != nil {
			errors.ApiError(w, http.StatusInternalServerError, "Failed To Add new User in database! \n"+result.Error.Error())
			return
		}

		payload := utils.Payload{
			Email: domainUser.Email,
		}

		token, err := utils.GenerateJwtToken(payload)
		if err != nil {
			errors.ApiError(w, http.StatusInternalServerError, "Failed To Generate New JWT Token!")
			return
		}

		utils.RespondWithJSON(w, SignupResponseOutput{
			Token: token,
			User:  signupResponse,
		})
	}
}
