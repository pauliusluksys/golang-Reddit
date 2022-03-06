package middlewares

import (
	"github.com/hashicorp/go-hclog"
	"github.com/pauliusluksys/golang-Reddit/services"
	"github.com/pauliusluksys/golang-Reddit/utils"
	"golang.org/x/net/context"
	"net/http"
)

type AuthMiddleware struct {
	AuthService services.AuthService
	Logger      hclog.Logger
}

func NewAuthMiddleware(authService services.AuthService, logger hclog.Logger) AuthMiddleware {
	return AuthMiddleware{AuthService: authService, Logger: logger}
}
func (s AuthMiddleware) CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Debug("logger is now fully functional horray!")
		authHeader := r.Header.Get("Authorization")
		user, err := s.AuthService.CheckAuth(authHeader)
		if err != nil {
			utils.RespondWithJSON(w, err.Code, err.Message)
		}
		ctx := context.WithValue(r.Context(), "user", user.FirstName+" "+user.LastName)
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
