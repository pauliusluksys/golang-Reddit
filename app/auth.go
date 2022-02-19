package app

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pauliusluksys/golang-Reddit/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users map[string][]byte = make(map[string][]byte)
var idxUsers int = 0

//getTokenUserPassword returns a jwt token for a user if the //password is ok
func getTokenUserPassword(w http.ResponseWriter, r *http.Request) {
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "cannot decode username/password struct", http.StatusBadRequest)
		return
	}
	//here I have a user!
	//Now check if exists
	passwordHash, found := users[u.Username]
	if !found {
		http.Error(w, "Cannot find the username", http.StatusNotFound)
	}
	err = bcrypt.CompareHashAndPassword(passwordHash, []byte(u.Password))
	if err != nil {
		return
	}
	token, err := createToken(u.Username)
	if err != nil {
		http.Error(w, "Cannot create token", http.StatusInternalServerError)
		return
	}
	sendJSONResponse(w, struct {
		Token string `json:"token"`
	}{token})
}
func createUser(w http.ResponseWriter, r *http.Request) {
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Cannot decode request", http.StatusBadRequest)
		return
	}
	if _, found := users[u.Username]; found {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}
	//If I'm here-> add user and return a token
	value, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	users[u.Username] = value
	token, err := createToken(u.Username)
	if err != nil {
		http.Error(w, "Cannot create token", http.StatusInternalServerError)
		return
	}
	sendJSONResponse(w, struct {
		Token string `json:"token"`
	}{token})
}
func createToken(username string) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"] = username
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	secret := utils.GetSecret()
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

type Username struct {
	username string
}

func isUsernameContextOk(username string, r *http.Request) bool {
	var usr Username
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&usr)
	if err != nil {
		fmt.Println("error: during username decoding: " + err.Error())
		return false
	}
	if usr.username != username {
		return false
	}
	return true
}
