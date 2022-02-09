package migrations

import (
	"fmt"
	"github.com/pauliusluksys/golang-Reddit/domain"
)

func PostMigration() {
	db := domain.GormDbConnections()
	err := db.AutoMigrate(&domain.PostGorm{})
	if err != nil {
		fmt.Println("Error while trying to migrate posts: " + err.Error())
	}
}
