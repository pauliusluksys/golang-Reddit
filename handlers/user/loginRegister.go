package userHandler

import (
	"encoding/json"
	"github.com/hashicorp/go-hclog"
	dtoUser "github.com/pauliusluksys/golang-Reddit/dto/user"
	"github.com/pauliusluksys/golang-Reddit/services"
	"github.com/pauliusluksys/golang-Reddit/utils"
	"net/http"
)

type UserHandler struct {
	UserService services.UserService
	Logger      hclog.Logger
}

var errors = utils.CustomError{}

type SignupResponseOutput struct {
	User  dtoUser.SignupResponse
	Token string
}
type LoginResponseOutput struct {
	User  dtoUser.LoginResponse
	Token string
}

func NewUserHandler(service services.UserService, logger hclog.Logger) UserHandler {
	return UserHandler{UserService: service, Logger: logger}
}

func (h UserHandler) UserLogin(w http.ResponseWriter, r *http.Request) {

	loginRequest := dtoUser.LoginRequest{}

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, err.Error())
	}
	loginResponse, cErr := h.UserService.Login(loginRequest)
	if cErr != nil {
		utils.RespondWithJSON(w, cErr.Code, cErr.Message)
	}

	//w.Header().Set("Authorization", "Bearer "+token)
	utils.RespondWithJSON(w, http.StatusOK, loginResponse)
}

func (h UserHandler) UserSignup(w http.ResponseWriter, r *http.Request) {
	signupRequest := dtoUser.SignupRequest{}
	err := json.NewDecoder(r.Body).Decode(&signupRequest)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "was not able to decode request")
	}
	signupResponse, cErr := h.UserService.Signup(signupRequest)
	if cErr != nil {
		utils.RespondWithJSON(w, cErr.Code, cErr.Message)
	}

	utils.RespondWithJSON(w, http.StatusOK, signupResponse)
}
