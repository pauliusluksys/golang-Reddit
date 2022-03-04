package middlewares

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/pauliusluksys/golang-Reddit/domain"
	"github.com/pauliusluksys/golang-Reddit/utils"
	"golang.org/x/net/context"
	"net/http"
	"strings"
)

var errors = utils.CustomError{}

type AuthService struct {
	logger  hclog.Logger
	configs *utils.Configurations
}

// NewAuthService returns a new instance of the auth service
func NewAuthService(logger hclog.Logger, configs *utils.Configurations) *AuthService {
	return &AuthService{logger, configs}
}
func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	//fmt.Println("wtf")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := domain.GormDbConnections()
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		fmt.Println("bearer token:    ", bearerToken)
		fmt.Println()
		if len(bearerToken) < 2 {
			errors.ApiError(w, http.StatusForbidden, "Token not provided!")
			return
		}

		token := bearerToken[1]
		fmt.Println(token)
		claims, err := utils.VerifyJwtToken(token)
		if err != nil {
			errors.ApiError(w, http.StatusForbidden, err.Error())
			return
		}
		user := domain.UserGorm{}
		if results := db.Where("email = ?", claims.Email).First(&user); results.Error != nil || results.RowsAffected < 1 {
			http.Error(w, "Unauthorized, user not exists", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "user_email", claims.Email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

//func isUsernameContextOk(username string, r *http.Request) bool {
//	usernameCtx := r.Context()
//
//	if usernameCtx.Value("user_email") != username {
//		return false
//	}
//	return true
//}
