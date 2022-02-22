package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) string {
	bPass := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bPass, bcrypt.MinCost)
	if err != nil {
		fmt.Println("error whilst comparing bcrypt: ", err)
	}
	return string(hash)
}
func DoPasswordsMatch(hashedPassword, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(currPassword))
	return err == nil
}
