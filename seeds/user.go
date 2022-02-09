package seeds

import (
	"github.com/bxcodec/faker/v3"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func (s Seed) UserSeed() {
	password := "password"
	bPass := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bPass, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < 100; i++ {
		//prepare the statement
		stmt, _ := s.db.Prepare(`INSERT INTO users(first_name,last_name,email,password,created_at) VALUES (?,?,?,?,?)`)
		// execute query

		_, err := stmt.Exec(faker.FirstName(), faker.LastName, faker.Email(), string(hash), faker.Date())
		if err != nil {
			panic(err)
		}
	}
}
